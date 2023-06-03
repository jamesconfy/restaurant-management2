package middleware

import (
	"fmt"
	"restaurant-management/internal/response"
	"restaurant-management/internal/se"
	"restaurant-management/internal/service"
	"strings"

	"github.com/casbin/casbin/v2"
	"github.com/gin-gonic/gin"
)

type JWT interface {
	CheckJWT() gin.HandlerFunc
}

type authMiddleWare struct {
	authSrv service.AuthService
	cashbin *casbin.Enforcer
}

func (a *authMiddleWare) CheckJWT() gin.HandlerFunc {
	return func(c *gin.Context) {
		authToken := a.getAuthorizationHeader(c)
		if authToken == "" {
			response.Error(c, *se.Unauthorized(nil, "token cannot be empty"))
			return
		}

		token, err := a.authSrv.Validate(authToken)
		if err != nil {
			response.Error(c, *se.Unauthorized(err, "invalid authorization token"))
			return
		}

		a.checkPolicy(c, token)

		c.Set("userId", token.Id)
		c.Set("role", token.Role)
		c.Next()
	}
}

func Authentication(authSrv service.AuthService, cashbin *casbin.Enforcer) JWT {
	return &authMiddleWare{authSrv: authSrv, cashbin: cashbin}
}

// Auxillary Function
func (a *authMiddleWare) getAuthorizationHeader(c *gin.Context) string {
	if a.isBrowser(c.Request.UserAgent()) {
		authtoken, _ := c.Cookie("Authorization")
		return authtoken
	}

	authHeader := strings.Replace(c.GetHeader("Authorization"), "Bearer ", "", 1)
	return authHeader
}

func (a *authMiddleWare) isBrowser(userAgent string) bool {
	switch {
	case strings.Contains(userAgent, "Mozilla"), strings.Contains(userAgent, "Chrome"), strings.Contains(userAgent, "Postman"), strings.Contains(userAgent, "Edge"), strings.Contains(userAgent, "Trident"):
		return true
	default:
		return false
	}
}

func (a *authMiddleWare) enforce(sub string, obj string, act string) (bool, error) {
	err := a.cashbin.LoadPolicy()
	if err != nil {
		return false, fmt.Errorf("failed to load policy from DB: %w", err)
	}

	ok, err := a.cashbin.Enforce(sub, obj, act)
	if err != nil {
		return false, err
	}

	return ok, nil
}

func (a *authMiddleWare) checkPolicy(c *gin.Context, token *service.Token) {
	ok, err := a.enforce(token.Id, c.Request.URL.Path, c.Request.Method)
	if err != nil {
		response.Error(c, *se.Internal(err, "Error when enforcing role based action"))
		return
	}

	if !ok {
		method := c.Request.Method
		description := "you are not authorized to change this resource"

		if method == "GET" {
			description = "you are not authorzied to view this resource"
		}

		response.Error(c, *se.Forbidden(description))
		return
	}
}
