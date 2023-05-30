package service_test

import (
	"restaurant-management/internal/forms"
	"testing"
)

func TestAddMenu(t *testing.T) {
	// Create a new user object
	menu := generateMenuForm()

	tests := []struct {
		name    string
		menu    *forms.Menu
		wantErr bool
	}{
		// TODO: Add test cases.
		{name: "Test with correct details", menu: menu, wantErr: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := menuSrv.Add(tt.menu)
			if (err != nil) != tt.wantErr {
				t.Errorf("menuSrv.Add() error = %v, wantErr %v", err, tt.wantErr)
			}

			_, err = menuSrv.Add(tt.menu)
			if (err != nil) != true {
				t.Errorf("menuSrv.Add2() error = %v, wantErr %v", err, true)
			}
		})
	}
}
