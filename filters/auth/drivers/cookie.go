package drivers

import (
	"github.com/gorilla/sessions"
	"morningo/config"
	"net/http"
)

var store = sessions.NewCookieStore([]byte(config.GetEnv().APP_SECRET))

type cacheAuthManager struct {
	name string
}

func NewCacheAuthDriver() *cacheAuthManager {
	return &cacheAuthManager{
		name: config.GetCookieConfig().NAME,
	}
}

func (cache *cacheAuthManager) Check(http *http.Request) bool {
	// read cookie
	session, err := store.Get(http, cache.name)
	if err != nil {
		return false
	}
	if session == nil {
		return false
	}
	if session.Values == nil {
		return false
	}
	if session.Values["id"] == nil {
		return false
	}
	return true
}

func (cache *cacheAuthManager) User(http *http.Request) interface{} {
	// get model user
	session, err := store.Get(http, cache.name)
	if err != nil {
		return session.Values
	}
	return session.Values
}

func (cache *cacheAuthManager) Login(http *http.Request, w http.ResponseWriter, user map[string]interface{}) interface{} {
	// write cookie
	session, err := store.Get(http, cache.name)
	if err != nil {
		return false
	}
	session.Values["id"] = user["id"]
	session.Save(http, w)
	return true
}

func (cache *cacheAuthManager) Logout(http *http.Request, w http.ResponseWriter) bool {
	// del cookie
	session, err := store.Get(http, cache.name)
	if err != nil {
		return false
	}
	session.Values["id"] = nil
	session.Save(http, w)
	return true
}
