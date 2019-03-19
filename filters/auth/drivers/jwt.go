package drivers

import (
	"encoding/json"
	"morningo/modules/log"
	jwtLib "github.com/dgrijalva/jwt-go"
	"github.com/dgrijalva/jwt-go/request"
	"morningo/config"
	"net/http"
	"strings"
	"time"
	"github.com/gin-gonic/gin"
	"errors"
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
	var keyFun jwtLib.Keyfunc
	keyFun = func(token *jwtLib.Token) (interface{}, error) {
		b := []byte(jwtAuth.secret)
		return b, nil
	}
	authJwtToken, err := request.ParseFromRequest(c.Request, request.OAuth2Extractor, keyFun)

	if err != nil {
		log.Println(err)
		return false
	}

	c.Set("User", map[string]interface{}{
		"token": authJwtToken,
	})

	return authJwtToken.Valid
}

// User is get the auth user from token string of the request header which
// contains the user ID. The token string must start with "Bearer "
func (jwtAuth *jwtAuthManager) User(c *gin.Context) interface{} {

	var jwtToken *jwtLib.Token
	if jwtUser, exist := c.Get("User"); !exist {
		tokenStr := strings.Replace(c.Request.Header.Get("Authorization"), "Bearer ", "", -1)
		if tokenStr == "" {
			return map[interface{}]interface{}{}
		}
		var err error
		jwtToken, err = jwtLib.Parse(tokenStr, func(token *jwtLib.Token) (interface{}, error) {
			b := []byte(jwtAuth.secret)
			return b, nil
		})
		if err != nil {
			panic(err)
		}
	} else {
		jwtToken = jwtUser.(map[string]interface{})["token"].(*jwtLib.Token)
	}

	if claims, ok := jwtToken.Claims.(jwtLib.MapClaims); ok && jwtToken.Valid {
		var user map[string]interface{}
		if err := json.Unmarshal([]byte(claims["user"].(string)), &user); err != nil {
			panic(err)
		}
		c.Set("User", map[string]interface{}{
			"token": jwtToken,
			"user":  user,
		})
		return user
	} else {
		panic(errors.New("decode jwt user claims fail"))
	}
}

func (jwtAuth *jwtAuthManager) Login(http *http.Request, w http.ResponseWriter, user map[string]interface{}) interface{} {

	token := jwtLib.New(jwtLib.GetSigningMethod(jwtAuth.alg))
	// Set some claims
	userStr, err := json.Marshal(user)
	token.Claims = jwtLib.MapClaims{
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

func (jwtAuth *jwtAuthManager) Logout(http *http.Request, w http.ResponseWriter) bool {
	// TODO: implement
	return true
}