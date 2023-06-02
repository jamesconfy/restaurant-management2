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

func TestAddFood_Admin(t *testing.T) {
	w := httptest.NewRecorder()

	food := generateFoodForm(nil)
	admin := generateAdminForm()
	form := generateLoginForm(admin)
	_ = createAndRegisterUser(admin)

	auth := loginUserAndGenerateAuth(form)

	obj, err := json.Marshal(food)
	if err != nil {
		panic(err)
	}

	req, _ := http.NewRequest("POST", utils.FoodPath, bytes.NewReader(obj))
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", auth)

	router.ServeHTTP(w, req)

	_, err = io.ReadAll(w.Body)
	if err != nil {
		panic(err)
	}

	assert.Equal(t, http.StatusOK, w.Code, "Status code should be the same")
}

func TestAddFood_User(t *testing.T) {
	w := httptest.NewRecorder()

	food := generateFoodForm(nil)
	user := generateUserForm()
	form := generateLoginForm(user)
	_ = createAndRegisterUser(user)

	auth := loginUserAndGenerateAuth(form)

	obj, err := json.Marshal(food)
	if err != nil {
		panic(err)
	}

	req, _ := http.NewRequest("POST", utils.FoodPath, bytes.NewReader(obj))
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", auth)

	router.ServeHTTP(w, req)

	_, err = io.ReadAll(w.Body)
	if err != nil {
		panic(err)
	}

	assert.Equal(t, http.StatusUnauthorized, w.Code, "Status code should be the same")
}

func TestGetFood(t *testing.T) {
	w := httptest.NewRecorder()

	food := createAndAddFood(nil, nil)

	url := fmt.Sprintf("%v/%v", utils.FoodPath, food.Id)

	req, _ := http.NewRequest("GET", url, nil)
	req.Header.Set("Content-Type", "application/json")

	router.ServeHTTP(w, req)

	_, err := io.ReadAll(w.Body)
	if err != nil {
		panic(err)
	}

	assert.Equal(t, http.StatusOK, w.Code, "Status code should be the same")
}

func TestGetAllFood(t *testing.T) {
	w := httptest.NewRecorder()

	for i := 0; i < 10; i++ {
		_ = createAndAddFood(nil, nil)
	}

	req, _ := http.NewRequest("GET", utils.FoodPath, nil)
	req.Header.Set("Content-Type", "application/json")

	router.ServeHTTP(w, req)

	_, err := io.ReadAll(w.Body)
	if err != nil {
		panic(err)
	}

	assert.Equal(t, http.StatusOK, w.Code, "Status code should be the same")
}

func TestEditFood_Admin(t *testing.T) {
	w := httptest.NewRecorder()

	menu := createAndAddMenu(nil)

	food := createAndAddFood(menu, nil)
	foo := generateFoodForm(menu)
	admin := generateAdminForm()
	form := generateLoginForm(admin)
	_ = createAndRegisterUser(admin)

	auth := loginUserAndGenerateAuth(form)

	obj, err := json.Marshal(foo)
	if err != nil {
		panic(err)
	}

	url := fmt.Sprintf("%v/%v", utils.FoodPath, food.Id)

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

func TestEditFood_User(t *testing.T) {
	w := httptest.NewRecorder()

	menu := createAndAddMenu(nil)

	food := createAndAddFood(menu, nil)
	foo := generateFoodForm(menu)
	user := generateUserForm()
	form := generateLoginForm(user)
	_ = createAndRegisterUser(user)

	auth := loginUserAndGenerateAuth(form)

	obj, err := json.Marshal(foo)
	if err != nil {
		panic(err)
	}

	url := fmt.Sprintf("%v/%v", utils.FoodPath, food.Id)

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

func TestDeleteFood_Admin(t *testing.T) {
	w := httptest.NewRecorder()

	menu := createAndAddMenu(nil)
	food := createAndAddFood(menu, nil)

	admin := generateAdminForm()
	form := generateLoginForm(admin)
	_ = createAndRegisterUser(admin)

	auth := loginUserAndGenerateAuth(form)

	url := fmt.Sprintf("%v/%v", utils.FoodPath, food.Id)

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

func TestDeleteFood_User(t *testing.T) {
	w := httptest.NewRecorder()

	menu := createAndAddMenu(nil)
	food := createAndAddFood(menu, nil)

	user := generateUserForm()
	form := generateLoginForm(user)
	_ = createAndRegisterUser(user)

	auth := loginUserAndGenerateAuth(form)

	url := fmt.Sprintf("%v/%v", utils.FoodPath, food.Id)

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
