package handler_test

import (
	"math/rand"
	"restaurant-management/internal/forms"
	"restaurant-management/internal/models"
)

func generateOrderForm(table *models.Table) *forms.Order {
	if table == nil {
		table = createAndAddTable(nil)
	}

	return &forms.Order{TableId: table.Id}
}

func createAndAddOrder(table *models.Table, order *forms.Order) *models.Order {
	if order == nil {
		order = generateOrderForm(table)
	}

	orde, err := orderSrv.Add(order)
	if err != nil {
		panic(err)
	}

	return orde
}

func generateEditOrderForm() *forms.EditOrder {
	return &forms.EditOrder{
		DeliveryId: rand.Intn(3) + 1,
		PaymentId:  rand.Intn(2) + 1,
	}
}
