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

func TestAddTable_Admin(t *testing.T) {
	w := httptest.NewRecorder()

	table := generateTableForm()
	admin := generateAdminForm()
	form := generateLoginForm(admin)
	_ = createAndRegisterUser(admin)

	auth := loginUserAndGenerateAuth(form)

	obj, err := json.Marshal(table)
	if err != nil {
		panic(err)
	}

	req, _ := http.NewRequest("POST", utils.TablePath, bytes.NewReader(obj))
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", auth)

	router.ServeHTTP(w, req)

	_, err = io.ReadAll(w.Body)
	if err != nil {
		panic(err)
	}

	assert.Equal(t, http.StatusOK, w.Code, "Status code should be the same")
}

func TestAddTable_User(t *testing.T) {
	w := httptest.NewRecorder()

	table := generateTableForm()
	user := generateUserForm()
	form := generateLoginForm(user)
	_ = createAndRegisterUser(user)

	auth := loginUserAndGenerateAuth(form)

	obj, err := json.Marshal(table)
	if err != nil {
		panic(err)
	}

	req, _ := http.NewRequest("POST", utils.TablePath, bytes.NewReader(obj))
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", auth)

	router.ServeHTTP(w, req)

	_, err = io.ReadAll(w.Body)
	if err != nil {
		panic(err)
	}

	assert.Equal(t, http.StatusForbidden, w.Code, "Status code should be the same")
}

func TestGetTable_Admin(t *testing.T) {
	w := httptest.NewRecorder()

	table := createAndAddTable(nil)
	admin := generateAdminForm()
	form := generateLoginForm(admin)
	_ = createAndRegisterUser(admin)

	auth := loginUserAndGenerateAuth(form)

	url := fmt.Sprintf("%v/%v", utils.TablePath, table.Id)

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

func TestGetAllTable_Admin(t *testing.T) {
	w := httptest.NewRecorder()

	admin := generateAdminForm()
	form := generateLoginForm(admin)
	_ = createAndRegisterUser(admin)

	auth := loginUserAndGenerateAuth(form)

	for i := 0; i < 10; i++ {
		_ = createAndAddTable(nil)
	}

	req, _ := http.NewRequest("GET", utils.TablePath, nil)
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", auth)

	router.ServeHTTP(w, req)

	_, err := io.ReadAll(w.Body)
	if err != nil {
		panic(err)
	}

	assert.Equal(t, http.StatusOK, w.Code, "Status code should be the same")
}

func TestGetTable_User(t *testing.T) {
	w := httptest.NewRecorder()

	table := createAndAddTable(nil)
	user := generateUserForm()
	form := generateLoginForm(user)
	_ = createAndRegisterUser(user)

	auth := loginUserAndGenerateAuth(form)

	url := fmt.Sprintf("%v/%v", utils.TablePath, table.Id)

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

func TestGetAllTable_User(t *testing.T) {
	w := httptest.NewRecorder()

	user := generateUserForm()
	form := generateLoginForm(user)
	_ = createAndRegisterUser(user)

	auth := loginUserAndGenerateAuth(form)

	for i := 0; i < 10; i++ {
		_ = createAndAddTable(nil)
	}

	req, _ := http.NewRequest("GET", utils.TablePath, nil)
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", auth)

	router.ServeHTTP(w, req)

	_, err := io.ReadAll(w.Body)
	if err != nil {
		panic(err)
	}

	assert.Equal(t, http.StatusOK, w.Code, "Status code should be the same")
}

func TestEditTable_Admin(t *testing.T) {
	w := httptest.NewRecorder()

	tabl := createAndAddTable(nil)

	table := generateEditTableForm()
	admin := generateAdminForm()
	form := generateLoginForm(admin)
	_ = createAndRegisterUser(admin)

	auth := loginUserAndGenerateAuth(form)

	obj, err := json.Marshal(table)
	if err != nil {
		panic(err)
	}

	url := fmt.Sprintf("%v/%v", utils.TablePath, tabl.Id)

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

func TestEditTable_User(t *testing.T) {
	w := httptest.NewRecorder()

	tabl := createAndAddTable(nil)

	table := generateEditTableForm()
	user := generateUserForm()
	form := generateLoginForm(user)
	_ = createAndRegisterUser(user)

	auth := loginUserAndGenerateAuth(form)

	obj, err := json.Marshal(table)
	if err != nil {
		panic(err)
	}

	url := fmt.Sprintf("%v/%v", utils.TablePath, tabl.Id)

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

func TestDeleteTable_Admin(t *testing.T) {
	w := httptest.NewRecorder()

	table := createAndAddTable(nil)
	admin := generateAdminForm()
	form := generateLoginForm(admin)
	_ = createAndRegisterUser(admin)

	auth := loginUserAndGenerateAuth(form)

	url := fmt.Sprintf("%v/%v", utils.TablePath, table.Id)

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

func TestDeleteTable_User(t *testing.T) {
	w := httptest.NewRecorder()

	table := createAndAddTable(nil)
	user := generateUserForm()
	form := generateLoginForm(user)
	_ = createAndRegisterUser(user)

	auth := loginUserAndGenerateAuth(form)

	url := fmt.Sprintf("%v/%v", utils.TablePath, table.Id)

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
