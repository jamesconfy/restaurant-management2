package service_test

import (
	"restaurant-management/internal/forms"
	"testing"
)

func TestAddTable(t *testing.T) {
	// Create a new user object
	table := generateTableForm()

	tests := []struct {
		name    string
		table   *forms.Table
		wantErr bool
	}{
		// TODO: Add test cases.
		{name: "Test with correct details", table: table, wantErr: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := tableSrv.Add(tt.table)
			if (err != nil) != tt.wantErr {
				t.Errorf("tableSrv.Add() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
func TestGetTable(t *testing.T) {
	// Create a new table object
	table := createAndAddTable(nil)

	tests := []struct {
		name    string
		id      string
		wantErr bool
	}{
		// TODO: Add test cases.
		{name: "Test with correct details", id: table.Id, wantErr: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := tableSrv.Get(tt.id)
			if (err != nil) != tt.wantErr {
				t.Errorf("tableSrv.Get() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestGetAllTable_Admin(t *testing.T) {
	for i := 0; i < 10; i++ {
		_ = generateTableForm()
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
			_, err := tableSrv.GetAll("ADMIN")
			if (err != nil) != tt.wantErr {
				t.Errorf("tableSrv.GetAll() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestGetAllTable_User(t *testing.T) {
	for i := 0; i < 10; i++ {
		_ = generateTableForm()
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
			_, err := tableSrv.GetAll("USER")
			if (err != nil) != tt.wantErr {
				t.Errorf("tableSrv.GetAll() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestEditTable(t *testing.T) {
	// Create a new user object
	table := createAndAddTable(nil)

	tabl := generateEditTableForm()

	tests := []struct {
		name    string
		id      string
		table   *forms.EditTable
		wantErr bool
	}{
		// TODO: Add test cases.
		{name: "Test with correct details", id: table.Id, table: tabl, wantErr: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := tableSrv.Edit(tt.id, tt.table)
			if (err != nil) != tt.wantErr {
				t.Errorf("tableSrv.Edit() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestDeleteTable(t *testing.T) {
	// Create a new user object
	table := createAndAddTable(nil)

	tests := []struct {
		name    string
		id      string
		wantErr bool
	}{
		// TODO: Add test cases.
		{name: "Test with correct details", id: table.Id, wantErr: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := tableSrv.Delete(tt.id)
			if (err != nil) != tt.wantErr {
				t.Errorf("tableSrv.Delete() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
