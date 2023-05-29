package repo_test

import (
	"math/rand"

	"restaurant-management/internal/models"
)

func generateTable() *models.Table {
	return &models.Table{
		Seats: rand.Intn(100) + 1,
	}
}

func createAndAddTable(table *models.Table) *models.Table {
	if table == nil {
		table = generateTable()
	}

	table, err := ta.Add(table)
	if err != nil {
		panic(err)
	}

	return table
}
