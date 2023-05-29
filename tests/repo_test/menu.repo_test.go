package repo_test

import (
	"restaurant-management/internal/models"
	"testing"
)

func TestAddMenu(t *testing.T) {
	menu := generateMenu()

	tests := []struct {
		name    string
		menu    *models.Menu
		wantErr bool
	}{
		// TODO: Add test cases.
		{name: "Test with correct user id", menu: menu, wantErr: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := m.Add(tt.menu)

			if (err != nil) != tt.wantErr {
				t.Errorf("menuSql.Add() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
