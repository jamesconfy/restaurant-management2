package routes

import (
	"restaurant-management/cmd/handlers"
	"restaurant-management/cmd/middleware"
	"restaurant-management/internal/service"

	"github.com/casbin/casbin/v2"
	"github.com/gin-gonic/gin"
)

func MenuRoute(router *gin.RouterGroup, menuSrv service.MenuService, authSrv service.AuthService, cashbin *casbin.Enforcer) {
	handler := handlers.NewMenuHandler(menuSrv, cashbin)
	jwt := middleware.Authentication(authSrv, cashbin)

	menu := router.Group("/menus")
	{
		menu.GET("/:menuId", handler.Get)
		menu.GET("", handler.GetAll)
	}

	menu.Use(jwt.CheckJWT())
	{
		menu.POST("", handler.Add)
		menu.PATCH("/:menuId", handler.Edit)
		menu.DELETE("/:menuId", handler.Delete)
	}
}
