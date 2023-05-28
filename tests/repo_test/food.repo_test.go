package repo_test

import (
	"fmt"
	"restaurant-management/internal/models"
	"testing"
)

func TestAddFood(t *testing.T) {
	menu := createAndAddMenu(nil)
	food := generateFood(menu)

	tests := []struct {
		name    string
		food    *models.Food
		wantErr bool
	}{
		// TODO: Add test cases.
		{name: "Test with correct details", food: food, wantErr: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := f.Add(tt.food)

			if (err != nil) != tt.wantErr {
				t.Errorf("foodSql.Add() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestGetFood(t *testing.T) {
	menu := createAndAddMenu(nil)
	food := createAndAddFood(menu, nil)

	fmt.Println(food)

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
			_, err := f.Get(tt.id)

			if (err != nil) != tt.wantErr {
				t.Errorf("foodSql.Get() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestGetAllFood(t *testing.T) {
	menu := createAndAddMenu(nil)
	for i := 0; i < 10; i++ {
		_ = createAndAddFood(menu, nil)
	}

	tests := []struct {
		name string

		wantErr bool
	}{
		// TODO: Add test cases.
		{name: "Test with correct details", wantErr: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := f.GetAll()

			if (err != nil) != tt.wantErr {
				t.Errorf("foodSql.GetAll() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestEditFood(t *testing.T) {
	menu := createAndAddMenu(nil)
	food := createAndAddFood(menu, nil)

	food2 := generateFood(menu)

	tests := []struct {
		name    string
		id      string
		food    *models.Food
		wantErr bool
	}{
		// TODO: Add test cases.
		{name: "Test with correct details", id: food.Id, food: food2, wantErr: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := f.Edit(tt.id, tt.food)

			if (err != nil) != tt.wantErr {
				t.Errorf("foodSql.Edit() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
