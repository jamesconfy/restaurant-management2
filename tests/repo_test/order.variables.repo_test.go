package repo_test

import "restaurant-management/internal/models"

func generateOrder(table *models.Table) *models.Order {
	if table == nil {
		table = createAndAddTable(nil)
	}

	return &models.Order{
		TableId: table.Id,
	}
}

func createAndAddOrder(table *models.Table, order *models.Order) *models.Order {
	if order == nil {
		order = generateOrder(table)
	}

	order, err := o.Add(order)
	if err != nil {
		panic(err)
	}

	return order
}
