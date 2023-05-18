package routes

import (
	"restaurant-management/cmd/handlers"
	"restaurant-management/cmd/middleware"
	"restaurant-management/internal/service"

	"github.com/gin-gonic/gin"
)

func UserRoute(router *gin.RouterGroup, userSrv service.UserService, authSrv service.AuthService) {
	handler := handlers.NewUserHandler(userSrv)
	middleware := middleware.Authentication(authSrv)

	user := router.Group("/users")
	user.Use(middleware.CheckJWT())
	{
		user.GET("/:userId", handler.Get)
		user.GET("", handler.GetAll)
	}

	auth1, auth2 := router.Group("/auth"), router.Group("/auth")
	{
		auth1.POST("/register", handler.Create)
		auth1.POST("/login", handler.Login)
	}

	auth2.Use(middleware.CheckJWT())
	{
		auth2.POST("/logout", handler.Logout)
		auth2.DELETE("/clear", handler.ClearAuth)
	}
}
