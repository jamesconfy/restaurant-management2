package handler_test

import (
	"io"
	"net/http"
	"net/http/httptest"
	"restaurant-management/utils"
	"testing"

	"github.com/go-playground/assert/v2"
)

func TestHome(t *testing.T) {
	r := setupApp()
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", utils.BasePath, nil)

	r.ServeHTTP(w, req)

	_, err := io.ReadAll(w.Body)
	if err != nil {
		panic(err)
	}

	assert.Equal(t, http.StatusOK, w.Code)
}
