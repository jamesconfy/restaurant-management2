package service_test

import (
	"fmt"
	"restaurant-management/internal/forms"
	"testing"
)

func TestAddOrder(t *testing.T) {
	order := generateOrder(nil)

	tests := []struct {
		name    string
		order   *forms.Order
		wantErr bool
	}{
		// TODO: Add test cases.
		{name: "Test with correct details", order: order, wantErr: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := orderSrv.Add(tt.order)
			if (err != nil) != tt.wantErr {
				t.Errorf("orderSrv.Add() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestGetOrder(t *testing.T) {
	// Create a new user object
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
			_, err := orderSrv.Get(tt.id)
			if (err != nil) != tt.wantErr {
				t.Errorf("orderSrv.Get() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestGetAllOrder(t *testing.T) {
	// Create a new user object
	table := createAndAddTable(nil)
	for i := 0; i < 10; i++ {
		_ = createAndAddOrder(table, nil)
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
			_, err := orderSrv.GetAll()
			if (err != nil) != tt.wantErr {
				t.Errorf("orderSrv.GetAll() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestEditOrder(t *testing.T) {
	// Create a new user object

	order := createAndAddOrder(nil, nil)
	req := generateEditOrderForm()

	fmt.Println("Old Order: ", order)

	tests := []struct {
		name    string
		id      string
		req     *forms.EditOrder
		wantErr bool
	}{
		// TODO: Add test cases.
		{name: "Test with correct details", id: order.Id, req: req, wantErr: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			order, err := orderSrv.Edit(tt.id, tt.req)
			fmt.Println("New Order: ", order)
			if (err != nil) != tt.wantErr {
				t.Errorf("orderSrv.Edit() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestDeleteOrder(t *testing.T) {
	// Create a new user object
	table := createAndAddTable(nil)
	order := createAndAddOrder(table, nil)

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
			err := orderSrv.Delete(tt.id)
			if (err != nil) != tt.wantErr {
				t.Errorf("orderSrv.Delete() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
