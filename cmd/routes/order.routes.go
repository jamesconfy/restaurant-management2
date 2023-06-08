package routes

import (
	"restaurant-management/cmd/handlers"
	"restaurant-management/cmd/middleware"
	"restaurant-management/internal/service"

	"github.com/casbin/casbin/v2"
	"github.com/gin-gonic/gin"
)

func OrderRoutes(router *gin.RouterGroup, orderSrv service.OrderService, authSrv service.AuthService, casbin *casbin.Enforcer) {
	handler := handlers.NewOrderHandler(orderSrv, casbin)
	jwt := middleware.Authentication(authSrv, casbin)

	order := router.Group("/orders")
	order.Use(jwt.CheckJWT())
	{
		order.POST("", handler.Add)
		order.GET("", handler.GetAll)
		order.GET("/:orderId", handler.Get)
		order.DELETE("/:orderId", handler.Delete)
		order.PATCH("/:orderId", handler.Edit)
	}
}
