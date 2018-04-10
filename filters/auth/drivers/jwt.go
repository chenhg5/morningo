package drivers

import (
	"encoding/json"
	"fmt"
	jwt_lib "github.com/dgrijalva/jwt-go"
	"github.com/dgrijalva/jwt-go/request"
	"log"
	"morningo/config"
	"net/http"
	"strings"
	"time"
)

type jwtAuthManager struct {
	secret string
	exp    time.Duration
	alg    string
}

func NewJwtAuthDriver() *jwtAuthManager {
	return &jwtAuthManager{
		secret: config.GetJwtConfig().SECRET,
		exp:    config.GetJwtConfig().EXP,
		alg:    config.GetJwtConfig().ALG,
	}
}

func (jwtAuth *jwtAuthManager) Check(http *http.Request) bool {
	token := http.Header.Get("Authorization")
	token = strings.Replace(token, "Bearer ", "", -1)
	if token == "" {
		return false
	}
	var keyFun jwt_lib.Keyfunc
	keyFun = func(token *jwt_lib.Token) (interface{}, error) {
		b := ([]byte(jwtAuth.secret))
		return b, nil
	}
	_, err := request.ParseFromRequest(http, request.OAuth2Extractor, keyFun)

	if err != nil {
		fmt.Println(err)
		return false
	}
	return true
}

func (jwtAuth *jwtAuthManager) User(http *http.Request) interface{} {
	// get model user
	tokenStr := http.Header.Get("Authorization")
	tokenStr = strings.Replace(tokenStr, "Bearer ", "", -1)
	if tokenStr == "" {
		return map[interface{}]interface{}{}
	}
	jwtToken, err := jwt_lib.Parse(tokenStr, func(token *jwt_lib.Token) (interface{}, error) {
		b := ([]byte(jwtAuth.secret))
		return b, nil
	})

	if claims, ok := jwtToken.Claims.(jwt_lib.MapClaims); ok && jwtToken.Valid {
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

	token := jwt_lib.New(jwt_lib.GetSigningMethod(jwtAuth.alg))
	// Set some claims
	userStr, err := json.Marshal(user)
	log.Println(string(userStr))
	token.Claims = jwt_lib.MapClaims{
		"user": string(userStr),
		"exp":  time.Now().Add(jwtAuth.exp).Unix(),
	}
	// Sign and get the complete encoded token as a string
	tokenString, err := token.SignedString([]byte(jwtAuth.secret))
	if err != nil {
		return nil
	}

	return tokenString
}

func (cache *jwtAuthManager) Logout(http *http.Request, w http.ResponseWriter) bool {
	return true
}
