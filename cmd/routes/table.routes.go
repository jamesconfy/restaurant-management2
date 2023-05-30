package routes

import (
	"restaurant-management/cmd/handlers"
	"restaurant-management/cmd/middleware"
	"restaurant-management/internal/service"

	"github.com/casbin/casbin/v2"
	"github.com/gin-gonic/gin"
)

func TableRoute(router *gin.RouterGroup, tableSrv service.TableService, authSrv service.AuthService, cashbin *casbin.Enforcer) {
	handler := handlers.NewTableHandler(tableSrv, cashbin)
	jwt := middleware.Authentication(authSrv, cashbin)

	table := router.Group("/tables")
	table.Use(jwt.CheckJWT())
	{
		table.POST("", handler.Add)
		table.GET("/:tableId", handler.Get)
		table.GET("", handler.GetAll)
		table.PATCH("/:tableId", handler.Edit)
		table.DELETE("/:tableId", handler.Delete)
	}
}
