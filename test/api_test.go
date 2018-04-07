package test

import (
	"net/http/httptest"
	"net/http"
	"encoding/base64"
	"testing"
	"github.com/stretchr/testify/assert"
	"github.com/gin-gonic/gin"
)

func TestIndexApi(t *testing.T) {
	called := false
	accounts := gin.Accounts{"foo": "bar"}
	router := gin.New()
	router.Use(gin.BasicAuth(accounts))
	router.GET("/login", func(c *gin.Context) {
		called = true
		c.String(200, c.MustGet(gin.AuthUserKey).(string))
	})

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/login", nil)
	req.Header.Set("Authorization", "Basic "+base64.StdEncoding.EncodeToString([]byte("admin:password")))
	router.ServeHTTP(w, req)

	assert.False(t, called)
	assert.Equal(t, 401, w.Code)
	assert.Equal(t, "Basic realm=\"Authorization Required\"", w.HeaderMap.Get("WWW-Authenticate"))
}