package drivers

import (
	"net/http"
	jwt "github.com/dgrijalva/jwt-go"
	"morningo/config"
	"time"
	"strings"
	"github.com/dgrijalva/jwt-go/request"
	"fmt"
	"encoding/json"
	"log"
)

type jwtAuthManager struct {
	name string
	exp  int64
}

func NewJwtAuthDriver() *jwtAuthManager{
	return &jwtAuthManager{
		name: config.GetCookieConfig().NAME,
		exp: time.Now().Add(time.Hour * 1).Unix(),
	}
}

func (jwtAuth *jwtAuthManager) Check(http *http.Request) bool  {
	token := http.Header.Get("Authentication")
	token = strings.Replace(token, "Bearer ", "", -1)
	if token == "" {
		return false
	}
	var keyFun jwt.Keyfunc
	keyFun = func(token *jwt.Token) (interface{}, error) {
		b := ([]byte(jwtAuth.name))
		return b, nil
	}
	_, err := request.ParseFromRequest(http, request.OAuth2Extractor, keyFun)

	if err != nil {
		return false
	}
	return true
}

func (jwtAuth *jwtAuthManager) User(http *http.Request) interface{}  {
	// get model user
	tokenStr := http.Header.Get("Authorization")
	tokenStr = strings.Replace(tokenStr, "Bearer ", "", -1)
	if tokenStr == "" {
		return map[interface{}]interface{}{}
	}
	jwtToken, err := jwt.Parse(tokenStr, func(token *jwt.Token) (interface{}, error) {
		b := ([]byte(jwtAuth.name))
		return b, nil
	})

	if claims, ok := jwtToken.Claims.(jwt.MapClaims); ok && jwtToken.Valid {
		var user map[string]interface{}
		if err := json.Unmarshal([]byte(claims["user"].(string)), &user); err != nil {
			fmt.Println(err)
			return map[interface{}]interface{}{}
		}
		return user
	} else {
		fmt.Println(err)
		return map[interface{}]interface{}{}
	}

}

func (jwtAuth *jwtAuthManager) Login(http *http.Request, w http.ResponseWriter, user map[string]interface{}) interface{} {

	token := jwt.New(jwt.GetSigningMethod("HS256"))
	// Set some claims
	userStr, err := json.Marshal(user)
	log.Println(string(userStr))
	token.Claims = jwt.MapClaims{
		"user":  string(userStr),
		"exp": time.Now().Add(time.Hour * 1).Unix(),
	}
	// Sign and get the complete encoded token as a string
	tokenString, err := token.SignedString([]byte(jwtAuth.name))
	if err != nil {
		return nil
	}

	return tokenString
}

func (cache *jwtAuthManager) Logout(http *http.Request, w http.ResponseWriter) bool  {
	return true
}