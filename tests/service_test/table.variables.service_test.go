package service_test

import (
	"math/rand"
	"restaurant-management/internal/forms"
	"restaurant-management/internal/models"
)

func generateTableForm() *forms.Table {
	return &forms.Table{
		Seats: rand.Intn(50) + 1,
	}
}

func generateEditTableForm() *forms.EditTable {
	return &forms.EditTable{
		Booked: true,
		Seats:  rand.Intn(50) + 1,
	}
}

func createAndAddTable(table *forms.Table) *models.Table {
	if table == nil {
		table = generateTableForm()
	}

	tabl, err := tableSrv.Add(table)
	if err != nil {
		panic(err)
	}

	return tabl
}
