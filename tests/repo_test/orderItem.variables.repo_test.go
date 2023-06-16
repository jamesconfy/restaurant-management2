package repo_test

import (
	"math/rand"
	"restaurant-management/internal/models"
)

func generateOrderItem(order *models.Order, food *models.Food) *models.OrderItem {
	if order == nil {
		order = createAndAddOrder(nil, nil)
	}

	if food == nil {
		food = createAndAddFood(nil, nil)
	}

	return &models.OrderItem{
		FoodId:   food.Id,
		OrderId:  order.Id,
		Quantity: rand.Intn(100) + 1,
	}
}

func createAndAddOrderItem(order *models.Order, food *models.Food) *models.OrderItem {
	orderItem := generateOrderItem(order, food)

	orderItem, err := orderItemRepo.Add(orderItem)
	if err != nil {
		panic(err)
	}

	return orderItem
}
