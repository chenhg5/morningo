package drivers

import (
	"encoding/json" // TODO: encoding/json运用了反射，略慢，需要考虑改进为：
	"fmt"
	jwt_lib "github.com/dgrijalva/jwt-go"
	"github.com/dgrijalva/jwt-go/request"
	"log"
	"morningo/config"
	"net/http"
	"strings"
	"time"
	"github.com/gin-gonic/gin"
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

// Check the token of request header is valid or not.
func (jwtAuth *jwtAuthManager) Check(c *gin.Context) bool {
	token := c.Request.Header.Get("Authorization")
	token = strings.Replace(token, "Bearer ", "", -1)
	if token == "" {
		return false
	}
	var keyFun jwt_lib.Keyfunc
	keyFun = func(token *jwt_lib.Token) (interface{}, error) {
		b := ([]byte(jwtAuth.secret))
		return b, nil
	}
	authJwtToken, err := request.ParseFromRequest(c.Request, request.OAuth2Extractor, keyFun)

	if err != nil {
		fmt.Println(err)
		return false
	}

	c.Set("User", map[string]interface{}{
		"token" : authJwtToken,
	})

	return authJwtToken.Valid
}

// User is get the auth user from token string of the request header which
// contains the user ID. The token string must start with "Bearer "
func (jwtAuth *jwtAuthManager) User(c *gin.Context) interface{} {

	var jwtToken *jwt_lib.Token
	if jwtUser, exist := c.Get("User"); !exist {
		tokenStr := c.Request.Header.Get("Authorization")
		tokenStr = strings.Replace(tokenStr, "Bearer ", "", -1)
		if tokenStr == "" {
			return map[interface{}]interface{}{}
		}
		var err error
		jwtToken, err = jwt_lib.Parse(tokenStr, func(token *jwt_lib.Token) (interface{}, error) {
			b := ([]byte(jwtAuth.secret))
			return b, nil
		})
		if err != nil {
			fmt.Println(err)
			return map[interface{}]interface{}{}
		}
	} else {
		jwtToken = jwtUser.(map[string]interface{})["token"].(*jwt_lib.Token)
	}

	if claims, ok := jwtToken.Claims.(jwt_lib.MapClaims); ok && jwtToken.Valid {
		var user map[string]interface{}
		if err := json.Unmarshal([]byte(claims["user"].(string)), &user); err != nil {
			fmt.Println(err)
			return map[interface{}]interface{}{}
		}
		c.Set("User", map[string]interface{}{
			"token" : jwtToken,
			"user" : user,
		})
		return user
	} else {
		fmt.Println(ok)
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
