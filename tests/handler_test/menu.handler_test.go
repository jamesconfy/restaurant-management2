package handler_test

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"restaurant-management/cmd/handlers"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAddMenu_Admin(t *testing.T) {
	w := httptest.NewRecorder()

	menu := generateMenuForm()
	admin := generateAdminForm()
	form := generateLoginForm(admin)
	_ = createAndRegisterUser(admin)

	auth := loginUserAndGenerateAuth(form)

	obj, err := json.Marshal(menu)
	if err != nil {
		panic(err)
	}

	req, _ := http.NewRequest("POST", handlers.MenuPath, bytes.NewReader(obj))
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", auth)

	router.ServeHTTP(w, req)

	_, err = io.ReadAll(w.Body)
	if err != nil {
		panic(err)
	}

	assert.Equal(t, http.StatusOK, w.Code, "Status code should be the same")
}

func TestAddMenu_User(t *testing.T) {
	w := httptest.NewRecorder()

	menu := generateMenuForm()
	user := generateUserForm()
	form := generateLoginForm(user)
	_ = createAndRegisterUser(user)

	auth := loginUserAndGenerateAuth(form)

	obj, err := json.Marshal(menu)
	if err != nil {
		panic(err)
	}

	req, _ := http.NewRequest("POST", handlers.MenuPath, bytes.NewReader(obj))
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", auth)

	router.ServeHTTP(w, req)

	_, err = io.ReadAll(w.Body)
	if err != nil {
		panic(err)
	}

	assert.Equal(t, http.StatusUnauthorized, w.Code, "Status code should be the same")
}

func TestGetMenu(t *testing.T) {
	w := httptest.NewRecorder()

	menu := createAndAddMenu(nil)

	url := fmt.Sprintf("%v/%v", handlers.MenuPath, menu.Id)

	req, _ := http.NewRequest("GET", url, nil)
	req.Header.Set("Content-Type", "application/json")

	router.ServeHTTP(w, req)

	_, err := io.ReadAll(w.Body)
	if err != nil {
		panic(err)
	}

	assert.Equal(t, http.StatusOK, w.Code, "Status code should be the same")
}

func TestGetAllMenu(t *testing.T) {
	w := httptest.NewRecorder()

	for i := 0; i < 10; i++ {
		_ = createAndAddMenu(nil)
	}

	req, _ := http.NewRequest("GET", handlers.MenuPath, nil)
	req.Header.Set("Content-Type", "application/json")

	router.ServeHTTP(w, req)

	_, err := io.ReadAll(w.Body)
	if err != nil {
		panic(err)
	}

	assert.Equal(t, http.StatusOK, w.Code, "Status code should be the same")
}

func TestEditMenu_Admin(t *testing.T) {
	w := httptest.NewRecorder()

	men := createAndAddMenu(nil)

	menu := generateMenuForm()
	admin := generateAdminForm()
	form := generateLoginForm(admin)
	_ = createAndRegisterUser(admin)

	auth := loginUserAndGenerateAuth(form)

	obj, err := json.Marshal(menu)
	if err != nil {
		panic(err)
	}

	url := fmt.Sprintf("%v/%v", handlers.MenuPath, men.Id)

	req, _ := http.NewRequest("PATCH", url, bytes.NewReader(obj))
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", auth)

	router.ServeHTTP(w, req)

	_, err = io.ReadAll(w.Body)
	if err != nil {
		panic(err)
	}

	assert.Equal(t, http.StatusOK, w.Code, "Status code should be the same")
}

func TestEditMenu_User(t *testing.T) {
	w := httptest.NewRecorder()

	men := createAndAddMenu(nil)

	menu := generateMenuForm()
	user := generateUserForm()
	form := generateLoginForm(user)
	_ = createAndRegisterUser(user)

	auth := loginUserAndGenerateAuth(form)

	obj, err := json.Marshal(menu)
	if err != nil {
		panic(err)
	}

	url := fmt.Sprintf("%v/%v", handlers.MenuPath, men.Id)

	req, _ := http.NewRequest("PATCH", url, bytes.NewReader(obj))
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", auth)

	router.ServeHTTP(w, req)

	_, err = io.ReadAll(w.Body)
	if err != nil {
		panic(err)
	}

	assert.Equal(t, http.StatusUnauthorized, w.Code, "Status code should be the same")
}

func TestDeleteMenu_Admin(t *testing.T) {
	w := httptest.NewRecorder()

	menu := createAndAddMenu(nil)
	admin := generateAdminForm()
	form := generateLoginForm(admin)
	_ = createAndRegisterUser(admin)

	auth := loginUserAndGenerateAuth(form)

	url := fmt.Sprintf("%v/%v", handlers.MenuPath, menu.Id)

	req, _ := http.NewRequest("DELETE", url, nil)
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", auth)

	router.ServeHTTP(w, req)

	_, err := io.ReadAll(w.Body)
	if err != nil {
		panic(err)
	}

	assert.Equal(t, http.StatusOK, w.Code, "Status code should be the same")
}

func TestDeleteMenu_User(t *testing.T) {
	w := httptest.NewRecorder()

	menu := createAndAddMenu(nil)
	user := generateUserForm()
	form := generateLoginForm(user)
	_ = createAndRegisterUser(user)

	auth := loginUserAndGenerateAuth(form)

	url := fmt.Sprintf("%v/%v", handlers.MenuPath, menu.Id)

	req, _ := http.NewRequest("DELETE", url, nil)
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", auth)

	router.ServeHTTP(w, req)

	_, err := io.ReadAll(w.Body)
	if err != nil {
		panic(err)
	}

	assert.Equal(t, http.StatusUnauthorized, w.Code, "Status code should be the same")
}
