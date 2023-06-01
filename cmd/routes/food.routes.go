package routes

import (
	"restaurant-management/cmd/handlers"
	"restaurant-management/cmd/middleware"
	"restaurant-management/internal/service"

	"github.com/casbin/casbin/v2"
	"github.com/gin-gonic/gin"
)

func FoodRoutes(router *gin.RouterGroup, foodSrv service.FoodService, authSrv service.AuthService, cashbin *casbin.Enforcer) {
	handler := handlers.NewFoodHanlder(foodSrv, cashbin)
	jwt := middleware.Authentication(authSrv, cashbin)

	food := router.Group("/foods")
	{
		food.GET("", handler.GetAll)
		food.GET("/:foodId", handler.Get)
	}

	food.Use(jwt.CheckJWT())
	{
		food.POST("", handler.Add)
		food.DELETE("/:foodId", handler.Delete)
		food.PATCH("/:foodId", handler.Edit)
	}
}
