package middleware

import (
	"fmt"
	"net/http"
	"restaurant-management/internal/service"
	"restaurant-management/utils"
	"strings"

	"github.com/gin-gonic/gin"
)

type JWT interface {
	CheckJWT() gin.HandlerFunc
}

type authMiddleWare struct {
	authSrv service.AuthService
}

func (a *authMiddleWare) CheckJWT() gin.HandlerFunc {
	return func(c *gin.Context) {
		authToken := GetAuthorizationHeader(c)
		if authToken == "" {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"message": "Invalid Authorization Token: Token cannot be empty"})
			return
		}

		token, err := a.authSrv.Validate(authToken)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"message": fmt.Sprintf("Invalid Authorization Token: %v", err)})
			return
		}

		fmt.Println("Request: ", c.Request.RequestURI)
		ok, err := enforce(fmt.Sprintf("role::%v", strings.ToLower(token.Role)), c.Request.RequestURI, c.Request.Method)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"message": fmt.Sprintf("Error when enforcing role base action: %v", err)})
			return
		}

		if !ok {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"message": "Unauthorized, you are not allowed to view this resource"})
		}

		c.Set("userId", token.Id)
		c.Set("role", token.Role)
		c.Next()
	}
}

func GetAuthorizationHeader(c *gin.Context) string {
	if isBrowser(c.Request.UserAgent()) {
		authtoken, _ := c.Cookie("Authorization")
		return authtoken
	}

	authHeader := strings.Replace(c.GetHeader("Authorization"), "Bearer ", "", 1)
	return authHeader
}

func Authentication(authSrv service.AuthService) JWT {
	return &authMiddleWare{authSrv: authSrv}
}

func isBrowser(userAgent string) bool {
	switch {
	case strings.Contains(userAgent, "Mozilla"), strings.Contains(userAgent, "Chrome"), strings.Contains(userAgent, "Postman"), strings.Contains(userAgent, "Edge"), strings.Contains(userAgent, "Trident"):
		return true
	default:
		return false
	}
}

func enforce(sub string, obj string, act string) (bool, error) {
	utils.Enforcer.AddGroupingPolicy("sample-user01", "user::role")
	utils.Enforcer.AddPolicy("sample-user02", "sample-path1", "GET", "allow")
	utils.Enforcer.SavePolicy()

	err := utils.Enforcer.LoadPolicy()
	if err != nil {
		return false, fmt.Errorf("failed to load policy from DB: %w", err)
	}

	ok, err := utils.Enforcer.Enforce(sub, obj, act)
	if err != nil {
		return false, err
	}

	return ok, nil
}
