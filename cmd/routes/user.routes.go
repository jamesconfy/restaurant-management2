package routes

import (
	"restaurant-management/cmd/handlers"
	"restaurant-management/cmd/middleware"
	"restaurant-management/internal/service"

	"github.com/casbin/casbin/v2"
	"github.com/gin-gonic/gin"
)

func UserRoute(router *gin.RouterGroup, userSrv service.UserService, authSrv service.AuthService, cashbin *casbin.Enforcer) {
	handler := handlers.NewUserHandler(userSrv, cashbin)
	jwt := middleware.Authentication(authSrv, cashbin)

	user := router.Group("/users")
	user.Use(jwt.CheckJWT())
	{
		user.GET("/:userId", handler.Get)
		user.GET("", handler.GetAll)
		user.GET("/profile", handler.GetProfile)
		user.PATCH("/profile", handler.Edit)
		user.DELETE("/profile", handler.Delete)
	}

	auth := router.Group("/auth")
	{
		auth.POST("/register", handler.Create)
		auth.POST("/login", handler.Login)
	}

	auth.Use(jwt.CheckJWT())
	{
		auth.POST("/register/admin", handler.Create)
		auth.POST("/logout", handler.Logout)
		auth.DELETE("/clear", handler.ClearAuth)
	}
}
