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
