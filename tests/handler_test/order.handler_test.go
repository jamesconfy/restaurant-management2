package handler_test

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"restaurant-management/utils"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAddOrder_Admin(t *testing.T) {
	w := httptest.NewRecorder()

	order := generateOrderForm(nil)
	admin := generateAdminForm()
	form := generateLoginForm(admin)
	_ = createAndRegisterUser(admin)

	auth := loginUserAndGenerateAuth(form)

	obj, err := json.Marshal(order)
	if err != nil {
		panic(err)
	}

	req, _ := http.NewRequest("POST", utils.OrderPath, bytes.NewReader(obj))
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", auth)

	router.ServeHTTP(w, req)

	_, err = io.ReadAll(w.Body)
	if err != nil {
		panic(err)
	}

	assert.Equal(t, http.StatusOK, w.Code, "Status code should be the same")
}

func TestAddOrder_User(t *testing.T) {
	w := httptest.NewRecorder()

	order := generateOrderForm(nil)
	user := generateUserForm()
	form := generateLoginForm(user)
	_ = createAndRegisterUser(user)

	auth := loginUserAndGenerateAuth(form)

	obj, err := json.Marshal(order)
	if err != nil {
		panic(err)
	}

	req, _ := http.NewRequest("POST", utils.OrderPath, bytes.NewReader(obj))
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", auth)

	router.ServeHTTP(w, req)

	_, err = io.ReadAll(w.Body)
	if err != nil {
		panic(err)
	}

	assert.Equal(t, http.StatusForbidden, w.Code, "Status code should be the same")
}

func TestGetOrder_Admin(t *testing.T) {
	w := httptest.NewRecorder()

	order := createAndAddOrder(nil, nil)
	admin := generateAdminForm()
	form := generateLoginForm(admin)
	_ = createAndRegisterUser(admin)

	auth := loginUserAndGenerateAuth(form)

	url := fmt.Sprintf("%v/%v", utils.OrderPath, order.Id)

	req, _ := http.NewRequest("GET", url, nil)
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", auth)

	router.ServeHTTP(w, req)

	_, err := io.ReadAll(w.Body)
	if err != nil {
		panic(err)
	}

	assert.Equal(t, http.StatusOK, w.Code, "Status code should be the same")
}

func TestGetAllOrder_Admin(t *testing.T) {
	w := httptest.NewRecorder()

	admin := generateAdminForm()
	form := generateLoginForm(admin)
	_ = createAndRegisterUser(admin)

	auth := loginUserAndGenerateAuth(form)

	for i := 0; i < 10; i++ {
		_ = createAndAddTable(nil)
	}

	req, _ := http.NewRequest("GET", utils.OrderPath, nil)
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", auth)

	router.ServeHTTP(w, req)

	_, err := io.ReadAll(w.Body)
	if err != nil {
		panic(err)
	}

	assert.Equal(t, http.StatusOK, w.Code, "Status code should be the same")
}

func TestGetOrder_User(t *testing.T) {
	w := httptest.NewRecorder()

	order := createAndAddOrder(nil, nil)
	user := generateUserForm()
	form := generateLoginForm(user)
	_ = createAndRegisterUser(user)

	auth := loginUserAndGenerateAuth(form)

	url := fmt.Sprintf("%v/%v", utils.OrderPath, order.Id)

	req, _ := http.NewRequest("GET", url, nil)
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", auth)

	router.ServeHTTP(w, req)

	_, err := io.ReadAll(w.Body)
	if err != nil {
		panic(err)
	}

	assert.Equal(t, http.StatusForbidden, w.Code, "Status code should be the same")
}

func TestGetAllOrder_User(t *testing.T) {
	w := httptest.NewRecorder()

	user := generateUserForm()
	form := generateLoginForm(user)
	_ = createAndRegisterUser(user)

	auth := loginUserAndGenerateAuth(form)

	for i := 0; i < 10; i++ {
		_ = createAndAddOrder(nil, nil)
	}

	req, _ := http.NewRequest("GET", utils.OrderPath, nil)
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", auth)

	router.ServeHTTP(w, req)

	_, err := io.ReadAll(w.Body)
	if err != nil {
		panic(err)
	}

	assert.Equal(t, http.StatusForbidden, w.Code, "Status code should be the same")
}

func TestEditOrder_Admin(t *testing.T) {
	w := httptest.NewRecorder()

	orde := createAndAddOrder(nil, nil)

	order := generateEditOrderForm()
	admin := generateAdminForm()
	form := generateLoginForm(admin)
	_ = createAndRegisterUser(admin)

	auth := loginUserAndGenerateAuth(form)

	obj, err := json.Marshal(order)
	if err != nil {
		panic(err)
	}

	url := fmt.Sprintf("%v/%v", utils.OrderPath, orde.Id)

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

func TestEditOrder_User(t *testing.T) {
	w := httptest.NewRecorder()

	orde := createAndAddOrder(nil, nil)

	order := generateEditOrderForm()
	user := generateUserForm()
	form := generateLoginForm(user)
	_ = createAndRegisterUser(user)

	auth := loginUserAndGenerateAuth(form)

	obj, err := json.Marshal(order)
	if err != nil {
		panic(err)
	}

	url := fmt.Sprintf("%v/%v", utils.OrderPath, orde.Id)

	req, _ := http.NewRequest("PATCH", url, bytes.NewReader(obj))
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", auth)

	router.ServeHTTP(w, req)

	_, err = io.ReadAll(w.Body)
	if err != nil {
		panic(err)
	}

	assert.Equal(t, http.StatusForbidden, w.Code, "Status code should be the same")
}

func TestDeleteOrder_Admin(t *testing.T) {
	w := httptest.NewRecorder()

	order := createAndAddOrder(nil, nil)
	admin := generateAdminForm()
	form := generateLoginForm(admin)
	_ = createAndRegisterUser(admin)

	auth := loginUserAndGenerateAuth(form)

	url := fmt.Sprintf("%v/%v", utils.OrderPath, order.Id)

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

func TestDeleteOrder_User(t *testing.T) {
	w := httptest.NewRecorder()

	order := createAndAddOrder(nil, nil)
	user := generateUserForm()
	form := generateLoginForm(user)
	_ = createAndRegisterUser(user)

	auth := loginUserAndGenerateAuth(form)

	url := fmt.Sprintf("%v/%v", utils.OrderPath, order.Id)

	req, _ := http.NewRequest("DELETE", url, nil)
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", auth)

	router.ServeHTTP(w, req)

	_, err := io.ReadAll(w.Body)
	if err != nil {
		panic(err)
	}

	assert.Equal(t, http.StatusForbidden, w.Code, "Status code should be the same")
}
