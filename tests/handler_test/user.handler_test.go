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

func TestCreateUser(t *testing.T) {
	w := httptest.NewRecorder()

	user := generateUserForm()

	obj, err := json.Marshal(user)
	if err != nil {
		panic(err)
	}

	url := fmt.Sprintf("%v/%v", utils.AuthPath, "register")

	req, _ := http.NewRequest("POST", url, bytes.NewReader(obj))
	req.Header.Set("Content-type", "application/json")

	router.ServeHTTP(w, req)

	_, err = io.ReadAll(w.Body)
	if err != nil {
		panic(err)
	}

	assert.Equal(t, http.StatusOK, w.Code, "Status code should be the same")
}

func TestCreateAdmin(t *testing.T) {
	w := httptest.NewRecorder()

	user := generateUserForm()
	admin := generateAdminForm()
	form := generateLoginForm(admin)
	_ = createAndRegisterUser(admin)

	auth := loginUserAndGenerateAuth(form)

	obj, err := json.Marshal(user)
	if err != nil {
		panic(err)
	}

	url := fmt.Sprintf("%v/%v", utils.AuthPath, "register/admin")

	req, _ := http.NewRequest("POST", url, bytes.NewReader(obj))
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", auth)

	router.ServeHTTP(w, req)

	_, err = io.ReadAll(w.Body)
	if err != nil {
		panic(err)
	}

	assert.Equal(t, http.StatusOK, w.Code, "Status code should be the same")
}

func TestLoginUser(t *testing.T) {
	w := httptest.NewRecorder()

	user := generateUserForm()
	_ = createAndRegisterUser(user)
	user1 := generateLoginForm(user)

	obj, err := json.Marshal(user1)
	if err != nil {
		panic(err)
	}

	url := fmt.Sprintf("%v/%v", utils.AuthPath, "login")

	req, _ := http.NewRequest("POST", url, bytes.NewReader(obj))
	req.Header.Set("Content-type", "application/json")

	router.ServeHTTP(w, req)

	_, err = io.ReadAll(w.Body)
	if err != nil {
		panic(err)
	}

	assert.Equal(t, http.StatusOK, w.Code, "Status code should be the same")
}

func TestGetUser_User(t *testing.T) {

	w := httptest.NewRecorder()

	user := generateUserForm()
	form := generateLoginForm(user)
	user1 := createAndRegisterUser(user)

	auth := loginUserAndGenerateAuth(form)

	url := fmt.Sprintf("%v/%v", utils.UserPath, user1.Id)

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

func TestGetProfile_User(t *testing.T) {

	w := httptest.NewRecorder()

	user := generateUserForm()
	form := generateLoginForm(user)
	_ = createAndRegisterUser(user)

	auth := loginUserAndGenerateAuth(form)

	url := fmt.Sprintf("%v/%v", utils.UserPath, "profile")

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

func TestGetProfile_Admin(t *testing.T) {

	w := httptest.NewRecorder()

	admin := generateAdminForm()
	form := generateLoginForm(admin)
	_ = createAndRegisterUser(admin)

	auth := loginUserAndGenerateAuth(form)

	url := fmt.Sprintf("%v/%v", utils.UserPath, "profile")

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

func TestGetAllUsers_User(t *testing.T) {
	w := httptest.NewRecorder()

	for i := 0; i < 10; i++ {
		_ = createAndRegisterUser(nil)
	}

	user := generateUserForm()
	form := generateLoginForm(user)
	_ = createAndRegisterUser(user)

	auth := loginUserAndGenerateAuth(form)

	req, _ := http.NewRequest("GET", utils.UserPath, nil)
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", auth)

	router.ServeHTTP(w, req)

	_, err := io.ReadAll(w.Body)
	if err != nil {
		panic(err)
	}

	assert.Equal(t, http.StatusUnauthorized, w.Code, "Status code should be the same")
}

func TestGetUser_Admin(t *testing.T) {
	w := httptest.NewRecorder()

	admin := generateAdminForm()
	form := generateLoginForm(admin)
	_ = createAndRegisterUser(admin)

	user := createAndRegisterUser(nil)

	auth := loginUserAndGenerateAuth(form)

	url := fmt.Sprintf("%v/%v", utils.UserPath, user.Id)

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

func TestGetAllUsers_Admin(t *testing.T) {
	w := httptest.NewRecorder()

	for i := 0; i < 10; i++ {
		_ = createAndRegisterUser(nil)
	}

	admin := generateAdminForm()
	form := generateLoginForm(admin)
	_ = createAndRegisterUser(admin)

	auth := loginUserAndGenerateAuth(form)

	req, _ := http.NewRequest("GET", utils.UserPath, nil)
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", auth)

	router.ServeHTTP(w, req)

	_, err := io.ReadAll(w.Body)
	if err != nil {
		panic(err)
	}

	assert.Equal(t, http.StatusOK, w.Code, "Status code should be the same")
}

func TestEditUser_User(t *testing.T) {
	w := httptest.NewRecorder()

	user := generateUserForm()
	user1 := generateUserForm()
	form := generateLoginForm(user)
	_ = createAndRegisterUser(user)

	auth := loginUserAndGenerateAuth(form)

	obj, err := json.Marshal(user1)
	if err != nil {
		panic(err)
	}

	url := fmt.Sprintf("%v/%v", utils.UserPath, "profile")

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

func TestEditUser_Admin(t *testing.T) {
	w := httptest.NewRecorder()

	admin := generateAdminForm()
	admin1 := generateAdminForm()
	form := generateLoginForm(admin)
	_ = createAndRegisterUser(admin)

	auth := loginUserAndGenerateAuth(form)

	obj, err := json.Marshal(admin1)
	if err != nil {
		panic(err)
	}

	url := fmt.Sprintf("%v/%v", utils.UserPath, "profile")

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

func TestDeleteUser_User(t *testing.T) {
	w := httptest.NewRecorder()

	user := generateUserForm()
	form := generateLoginForm(user)
	_ = createAndRegisterUser(user)

	auth := loginUserAndGenerateAuth(form)

	url := fmt.Sprintf("%v/%v", utils.UserPath, "profile")

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

func TestDeleteUser_Admin(t *testing.T) {
	w := httptest.NewRecorder()

	admin := generateAdminForm()
	form := generateLoginForm(admin)
	_ = createAndRegisterUser(admin)

	auth := loginUserAndGenerateAuth(form)

	url := fmt.Sprintf("%v/%v", utils.UserPath, "profile")

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
