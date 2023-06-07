package repo_test

import (
	"fmt"
	"restaurant-management/internal/models"
	"testing"
)

func TestAddOrder(t *testing.T) {
	order := generateOrder(nil)

	tests := []struct {
		name    string
		order   *models.Order
		wantErr bool
	}{
		// TODO: Add test cases.
		{name: "Test with correct details", order: order, wantErr: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := o.Add(tt.order)
			if (err != nil) != tt.wantErr {
				t.Errorf("orderSql.Add() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestGetOrder(t *testing.T) {
	order := createAndAddOrder(nil, nil)

	tests := []struct {
		name    string
		id      string
		wantErr bool
	}{
		// TODO: Add test cases.
		{name: "Test with correct details", id: order.Id, wantErr: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := o.Get(tt.id)
			if (err != nil) != tt.wantErr {
				t.Errorf("orderSql.Get() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestGetAllOrder(t *testing.T) {
	for i := 0; i < 10; i++ {
		_ = createAndAddOrder(nil, nil)
	}

	tests := []struct {
		name    string
		wantErr bool
	}{
		// TODO: Add test cases.
		{name: "Test with correct details", wantErr: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			orders, err := o.GetAll()

			for _, order := range orders {
				fmt.Printf("OrderId: %v	TableId: %v\n", order.Id, order.TableId)
			}
			if (err != nil) != tt.wantErr {
				t.Errorf("order.GetAll() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

// func TestGetAllTable_User(t *testing.T) {
// 	for i := 0; i < 10; i++ {
// 		_ = createAndAddTable(nil)
// 	}

// 	tests := []struct {
// 		name    string
// 		wantErr bool
// 	}{
// 		// TODO: Add test cases.
// 		{name: "Test with correct details", wantErr: false},
// 	}
// 	for _, tt := range tests {
// 		t.Run(tt.name, func(t *testing.T) {
// 			_, err := ta.GetAll("USER")
// 			if (err != nil) != tt.wantErr {
// 				t.Errorf("tableSql.GetAll() error = %v, wantErr %v", err, tt.wantErr)
// 			}
// 		})
// 	}
// }

// func TestEditTable(t *testing.T) {
// 	table := createAndAddTable(nil)

// 	tabl := generateTable()
// 	tabl.Booked = true

// 	tests := []struct {
// 		name    string
// 		id      string
// 		table   *models.Table
// 		wantErr bool
// 	}{
// 		// TODO: Add test cases.
// 		{name: "Test with correct details", id: table.Id, table: tabl, wantErr: false},
// 	}
// 	for _, tt := range tests {
// 		t.Run(tt.name, func(t *testing.T) {
// 			_, err := ta.Edit(tt.id, tt.table)

// 			if (err != nil) != tt.wantErr {
// 				t.Errorf("tableSql.Edit() error = %v, wantErr %v", err, tt.wantErr)
// 			}
// 		})
// 	}
// }

// func TestDeleteTable(t *testing.T) {
// 	table := createAndAddTable(nil)

// 	tests := []struct {
// 		name    string
// 		id      string
// 		wantErr bool
// 	}{
// 		// TODO: Add test cases.
// 		{name: "Test with correct details", id: table.Id, wantErr: false},
// 	}
// 	for _, tt := range tests {
// 		t.Run(tt.name, func(t *testing.T) {
// 			err := ta.Delete(tt.id)
// 			if (err != nil) != tt.wantErr {
// 				t.Errorf("tableSql.GetAll() error = %v, wantErr %v", err, tt.wantErr)
// 			}
// 		})
// 	}
// }
