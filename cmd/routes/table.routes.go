package routes

import (
	"restaurant-management/cmd/handlers"
	"restaurant-management/cmd/middleware"
	"restaurant-management/internal/service"

	"github.com/gin-gonic/gin"
)

func TableRoute(router *gin.RouterGroup, tableSrv service.TableService, jwt middleware.JWT) {
	handler := handlers.NewTableHandler(tableSrv)

	table := router.Group("/tables")
	table.Use(jwt.CheckJWT())
	{
		table.POST("", handler.Add)
		table.GET("/:tableId", handler.Get)
		table.GET("", handler.GetAll)
		table.DELETE("/:tableId", handler.Delete)
	}
}
