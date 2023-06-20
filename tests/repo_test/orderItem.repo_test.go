package repo_test

import (
	"restaurant-management/internal/models"
	"testing"
)

func TestAddOrderItem(t *testing.T) {
	item := generateOrderItem(nil, nil)

	tests := []struct {
		name    string
		item    *models.OrderItem
		wantErr bool
	}{
		// TODO: Add test cases.
		{name: "Test with correct details", item: item, wantErr: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := orderItemRepo.Add(tt.item)
			if (err != nil) != tt.wantErr {
				t.Errorf("orderItemSql.Add() error = %v, wantErr %v", err, tt.wantErr)
			}

			_, err = orderItemRepo.Add(tt.item)
			if (err != nil) != tt.wantErr {
				t.Errorf("orderItemSql.Add() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestGetOrderItem(t *testing.T) {
	item := createAndAddOrderItem(nil, nil)

	tests := []struct {
		name    string
		id      string
		wantErr bool
	}{
		// TODO: Add test cases.
		{name: "Test with correct details", id: item.Id, wantErr: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := orderItemRepo.Get(tt.id)
			if (err != nil) != tt.wantErr {
				t.Errorf("orderItemSql.Get() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestGetAllOrderItems(t *testing.T) {
	for i := 0; i < 10; i++ {
		_ = createAndAddOrderItem(nil, nil)
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
			_, err := orderItemRepo.GetAll()
			if (err != nil) != tt.wantErr {
				t.Errorf("orderItemSql.GetAll() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
