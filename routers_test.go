package main

import (
	"net/http/httptest"
	"net/http"
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestIndexApi(t *testing.T) {
	router := initRouter()

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/api/index", nil)
	router.ServeHTTP(w, req)

	assert.Equal(t, 200, w.Code)
	// assert.Equal(t, "GO GO GO!", w.Body.String())
}