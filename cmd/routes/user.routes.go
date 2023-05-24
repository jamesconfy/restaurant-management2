package routes

import (
	"restaurant-management/cmd/handlers"
	"restaurant-management/cmd/middleware"
	"restaurant-management/internal/service"

	"github.com/gin-gonic/gin"
)

func UserRoute(router *gin.RouterGroup, userSrv service.UserService, jwt middleware.JWT) {
	handler := handlers.NewUserHandler(userSrv)

	user := router.Group("/users")
	user.Use(jwt.CheckJWT())
	{
		user.GET("/:userId", handler.Get)
		user.GET("", handler.GetAll)
	}

	auth1, auth2 := router.Group("/auth"), router.Group("/auth")
	{
		auth1.POST("/register", handler.Create)
		auth1.POST("/login", handler.Login)
	}

	auth2.Use(jwt.CheckJWT())
	{
		auth2.POST("/register/admin", handler.Create)
		auth2.POST("/logout", handler.Logout)
		auth2.DELETE("/clear", handler.ClearAuth)
	}
}
