package service_test

import (
	"math/rand"
	"restaurant-management/internal/forms"
	"restaurant-management/internal/models"
)

func generateOrder(table *models.Table) *forms.Order {
	if table == nil {
		table = createAndAddTable(nil)
	}

	return &forms.Order{
		TableId: table.Id,
	}
}

func generateEditOrderForm() *forms.EditOrder {
	return &forms.EditOrder{
		PaymentId:  rand.Intn(2) + 1,
		DeliveryId: rand.Intn(3) + 1,
	}
}

func createAndAddOrder(table *models.Table, order *forms.Order) *models.Order {
	if order == nil {
		order = generateOrder(table)
	}

	orde, err := orderSrv.Add(order)
	if err != nil {
		panic(err)
	}

	return orde
}
