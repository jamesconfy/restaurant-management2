package repo_test

import (
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
		{name: "Test with correct user id", food: food, wantErr: false},
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
