package service_test

import (
	"restaurant-management/internal/forms"
	"testing"
)

func TestAddFood(t *testing.T) {
	menu := createAndAddMenu(nil)
	food := generateFoodForm(menu)

	tests := []struct {
		name    string
		menuId  string
		food    *forms.Food
		wantErr bool
	}{
		// TODO: Add test cases.
		{name: "Test with correct details", food: food, wantErr: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := foodSrv.Add(tt.food)
			if (err != nil) != tt.wantErr {
				t.Errorf("foodSrv.Add() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestGetFood(t *testing.T) {
	menu := createAndAddMenu(nil)
	food := createAndAddFood(menu, nil)

	tests := []struct {
		name    string
		id      string
		wantErr bool
	}{
		// TODO: Add test cases.
		{name: "Test with correct details", id: food.Id, wantErr: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := foodSrv.Get(tt.id)
			if (err != nil) != tt.wantErr {
				t.Errorf("foodSrv.Get() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestGetAllFood(t *testing.T) {
	for i := 0; i < 10; i++ {
		_ = createAndAddFood(nil, nil)
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
			_, err := foodSrv.GetAll()
			if (err != nil) != tt.wantErr {
				t.Errorf("foodSrv.GetAll() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestEditFood(t *testing.T) {
	menu := createAndAddMenu(nil)
	food := createAndAddFood(menu, nil)

	foo := generateEditFoodForm(nil)

	tests := []struct {
		name    string
		id      string
		food    *forms.EditFood
		wantErr bool
	}{
		// TODO: Add test cases.
		{name: "Test with correct details", id: food.Id, food: foo, wantErr: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := foodSrv.Edit(tt.id, tt.food)
			if (err != nil) != tt.wantErr {
				t.Errorf("foodSrv.Delete() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestDeleteFood(t *testing.T) {
	menu := createAndAddMenu(nil)
	food := createAndAddFood(menu, nil)

	tests := []struct {
		name    string
		id      string
		wantErr bool
	}{
		// TODO: Add test cases.
		{name: "Test with correct details", id: food.Id, wantErr: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := foodSrv.Delete(tt.id)
			if (err != nil) != tt.wantErr {
				t.Errorf("foodSrv.Delete() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
