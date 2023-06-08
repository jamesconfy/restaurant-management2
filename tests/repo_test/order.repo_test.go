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
			_, err := orderRepo.Add(tt.order)
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
			_, err := orderRepo.Get(tt.id)

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
			orders, err := orderRepo.GetAll()

			for _, order := range orders {
				fmt.Printf("OrderId: %v	TableId: %v\n", order.Id, order.TableId)
			}
			if (err != nil) != tt.wantErr {
				t.Errorf("order.GetAll() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestEditOrder(t *testing.T) {
	table := createAndAddTable(nil)
	order := createAndAddOrder(table, nil)

	orde := generateOrder(table)

	tests := []struct {
		name    string
		id      string
		order   *models.Order
		wantErr bool
	}{
		// TODO: Add test cases.
		{name: "Test with correct details", id: order.Id, order: orde, wantErr: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := orderRepo.Edit(tt.id, tt.order)

			if (err != nil) != tt.wantErr {
				t.Errorf("orderSql.Edit() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestDeleteOrder(t *testing.T) {
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
			err := orderRepo.Delete(tt.id)
			if (err != nil) != tt.wantErr {
				t.Errorf("orderSql.GetAll() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
