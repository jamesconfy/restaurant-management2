package repo_test

import (
	"restaurant-management/internal/models"
	"testing"
)

func TestNameAndCategory(t *testing.T) {
	menu := createAndAddMenu(nil)

	tests := []struct {
		name         string
		menuName     string
		menuCategory string
		wantErr      bool
		wantOk       bool
	}{
		// TODO: Add test cases.
		{name: "Test with correct details", menuName: menu.Name, menuCategory: menu.Category, wantErr: false, wantOk: true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ok, err := m.Check(tt.menuName, tt.menuCategory)
			if (err != nil) != tt.wantErr {
				t.Errorf("menuSql.Check() error = %v, wantErr %v", err, tt.wantErr)
			}

			if !ok {
				t.Errorf("menuSql.Check() ok = %v, wantOk %v", ok, tt.wantOk)
			}
		})
	}
}

func TestMenuExists(t *testing.T) {
	menu := createAndAddMenu(nil)

	tests := []struct {
		name    string
		id      string
		wantErr bool
		wantOk  bool
	}{
		// TODO: Add test cases.
		{name: "Test with correct details", id: menu.Id, wantErr: false, wantOk: true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ok, err := m.MenuExists(tt.id)
			if (err != nil) != tt.wantErr {
				t.Errorf("menuSql.CheckMenuExists() error = %v, wantErr %v", err, tt.wantErr)
			}

			if !ok {
				t.Errorf("menuSql.CheckMenuExists() ok = %v, wantOk %v", ok, tt.wantOk)
			}
		})
	}
}

func TestAddMenu(t *testing.T) {
	menu := generateMenu()

	tests := []struct {
		name    string
		menu    *models.Menu
		wantErr bool
	}{
		// TODO: Add test cases.
		{name: "Test with correct details", menu: menu, wantErr: false},
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

func TestGetMenu(t *testing.T) {
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
			_, err := m.Get(tt.id)

			if (err != nil) != tt.wantErr {
				t.Errorf("menuSql.Get() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestEditMenu(t *testing.T) {
	menu := createAndAddMenu(nil)
	editMenu := generateMenu()

	tests := []struct {
		name    string
		id      string
		menu    *models.Menu
		wantErr bool
	}{
		// TODO: Add test cases.
		{name: "Test with correct details", id: menu.Id, menu: editMenu, wantErr: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := m.Edit(tt.id, tt.menu)

			if (err != nil) != tt.wantErr {
				t.Errorf("menuSql.Edit() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestDeleteMenu(t *testing.T) {
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
			err := m.Delete(tt.id)

			if (err != nil) != tt.wantErr {
				t.Errorf("menuSql.Delete() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestGetAllMenu(t *testing.T) {
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
			_, err := m.GetAll()

			if (err != nil) != tt.wantErr {
				t.Errorf("menuSql.GetAll() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
