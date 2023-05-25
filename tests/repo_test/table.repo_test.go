package repo_test

import (
	"fmt"
	"restaurant-management/internal/models"
	"testing"
)

func TestAddTable(t *testing.T) {
	table := generateTable()

	tests := []struct {
		name    string
		table   *models.Table
		wantErr bool
	}{
		// TODO: Add test cases.
		{name: "Test with correct details", table: table, wantErr: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := ta.Add(tt.table)
			if (err != nil) != tt.wantErr {
				t.Errorf("tableSql.Add() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestGetTable(t *testing.T) {
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
			table, err := ta.Get(tt.id)
			fmt.Println("Table: ", table)
			if (err != nil) != tt.wantErr {
				t.Errorf("tableSql.Get() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestGetAllTable(t *testing.T) {
	for i := 0; i < 10; i++ {
		_ = createAndAddTable(nil)
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
			tables, err := ta.GetAll()

			for _, table := range tables {
				fmt.Println("Table: ", table)
			}

			if (err != nil) != tt.wantErr {
				t.Errorf("tableSql.GetAll() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestDeleteTable(t *testing.T) {
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
			err := ta.Delete(tt.id)
			if (err != nil) != tt.wantErr {
				t.Errorf("tableSql.GetAll() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}