package repo_test

import (
	"math/rand"
	"restaurant-management/internal/models"
)

func generateOrder(table *models.Table) *models.Order {
	if table == nil {
		table = createAndAddTable(nil)
	}

	return &models.Order{
		TableId:    table.Id,
		DeliveryId: rand.Intn(3) + 1,
		PaymentId:  rand.Intn(2) + 1,
	}
}

func createAndAddOrder(table *models.Table, order *models.Order) *models.Order {
	if order == nil {
		order = generateOrder(table)
	}

	order, err := orderRepo.Add(order)
	if err != nil {
		panic(err)
	}

	return order
}
