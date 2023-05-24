package middleware

import (
	"fmt"
	"net/http"
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
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"message": "Invalid Authorization Token: Token cannot be empty"})
			return
		}

		token, err := a.authSrv.Validate(authToken)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"message": fmt.Sprintf("Invalid Authorization Token: %v", err)})
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
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"message": fmt.Sprintf("Error when enforcing role base action: %v", err)})
		return
	}

	if !ok {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"message": "Unauthorized, you are not allowed to view this resource"})
		return
	}
}
