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

func TestGetMenu(t *testing.T) {
	// Create a new user object
	menu := createAndAddMenu(nil)

	tests := []struct {
		name    string
		id      string
		wantErr bool
	}{
		// TODO: Add test cases.
		{name: "Test with correct details", id: menu.Id, wantErr: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := menuSrv.Get(tt.id)
			if (err != nil) != tt.wantErr {
				t.Errorf("menuSrv.Get() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestGetAllMenu(t *testing.T) {
	// Create a new user object
	for i := 0; i < 10; i++ {
		_ = createAndAddMenu(nil)
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
			_, err := menuSrv.GetAll()
			if (err != nil) != tt.wantErr {
				t.Errorf("menuSrv.GetAll() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestEditMenu(t *testing.T) {
	// Create a new user object
	menu := createAndAddMenu(nil)
	for i := 0; i < 5; i++ {
		_ = createAndAddFood(menu, nil)

	}
	req := generateEditMenuForm()

	tests := []struct {
		name    string
		id      string
		req     *forms.EditMenu
		wantErr bool
	}{
		// TODO: Add test cases.
		{name: "Test with correct details", id: menu.Id, req: req, wantErr: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := menuSrv.Edit(tt.id, tt.req)
			if (err != nil) != tt.wantErr {
				t.Errorf("menuSrv.Edit() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestDeleteMenu(t *testing.T) {
	// Create a new user object
	menu := createAndAddMenu(nil)

	tests := []struct {
		name    string
		id      string
		wantErr bool
	}{
		// TODO: Add test cases.
		{name: "Test with correct details", id: menu.Id, wantErr: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := menuSrv.Delete(tt.id)
			if (err != nil) != tt.wantErr {
				t.Errorf("menuSrv.Delete() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
