package service_test

import (
	"restaurant-management/internal/forms"
	"testing"
)

func TestAddUser(t *testing.T) {
	// Create a new user object
	user := generateUserForm()

	tests := []struct {
		name    string
		user    *forms.User
		wantErr bool
	}{
		// TODO: Add test cases.
		{name: "Test with correct details", user: user, wantErr: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := userSrv.Add(tt.user)
			if (err != nil) != tt.wantErr {
				t.Errorf("userSrv.Add() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestLoginUser(t *testing.T) {
	// Create a new user object
	userForm := generateUserForm()
	form := &forms.Login{Email: userForm.Email, Password: userForm.Password}
	_ = createAndRegisterUser(userForm)

	tests := []struct {
		name    string
		form    *forms.Login
		wantErr bool
	}{
		// TODO: Add test cases.
		{name: "Test with correct details", form: form, wantErr: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := userSrv.Login(tt.form)
			if (err != nil) != tt.wantErr {
				t.Errorf("userSrv.Login() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestGetUser(t *testing.T) {
	// Create a new user object
	user := createAndRegisterUser(nil)

	tests := []struct {
		name    string
		id      string
		wantErr bool
	}{
		// TODO: Add test cases.
		{name: "Test with correct details", id: user.Id, wantErr: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := userSrv.Get(tt.id)
			if (err != nil) != tt.wantErr {
				t.Errorf("userSrv.Get() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestGetAll(t *testing.T) {
	// Create a new user object
	for i := 0; i < 10; i++ {
		_ = createAndRegisterUser(nil)
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
			_, err := userSrv.GetAll()
			if (err != nil) != tt.wantErr {
				t.Errorf("userSrv.GetAll() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestEditUser(t *testing.T) {
	// Create a new user object
	user := createAndRegisterUser(nil)
	edit := generateEditUserForm()

	tests := []struct {
		name    string
		id      string
		edit    *forms.EditUser
		wantErr bool
	}{
		// TODO: Add test cases.
		{name: "Test with correct details", id: user.Id, edit: edit, wantErr: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := userSrv.Edit(tt.id, tt.edit)
			if (err != nil) != tt.wantErr {
				t.Errorf("userSrv.Edit() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestDeleteUser(t *testing.T) {
	// Create a new user object
	user := createAndRegisterUser(nil)

	tests := []struct {
		name    string
		id      string
		wantErr bool
	}{
		// TODO: Add test cases.
		{name: "Test with correct details", id: user.Id, wantErr: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := userSrv.Delete(tt.id)
			if (err != nil) != tt.wantErr {
				t.Errorf("userSrv.Delete() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestDeleteAuth(t *testing.T) {
	// Create a new user object
	user := createAndRegisterUser(nil)
	auth := createUserAuth(user)

	tests := []struct {
		name        string
		id          string
		accessToken string
		wantErr     bool
	}{
		// TODO: Add test cases.
		{name: "Test with correct details", id: user.Id, accessToken: auth.AccessToken, wantErr: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := userSrv.DeleteAuth(tt.id, tt.accessToken)
			if (err != nil) != tt.wantErr {
				t.Errorf("userSrv.DeleteAuth() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestClearAuth(t *testing.T) {
	// Create a new user object
	user := createAndRegisterUser(nil)
	auth := createUserAuth(user)

	tests := []struct {
		name        string
		id          string
		accessToken string
		wantErr     bool
	}{
		// TODO: Add test cases.
		{name: "Test with correct details", id: user.Id, accessToken: auth.AccessToken, wantErr: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := userSrv.ClearAuth(tt.id, tt.accessToken)
			if (err != nil) != tt.wantErr {
				t.Errorf("userSrv.ClearAuth() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
