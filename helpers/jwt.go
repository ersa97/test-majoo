package helpers

import (
	"net/http"
	"time"

	"github.com/dgrijalva/jwt-go"
)

var JWTKeys = "test-case123"

func GenerateToken(userid int, name, username string) string {

	token := jwt.New(jwt.SigningMethodHS256)
	value := token.Claims.(jwt.MapClaims)

	value["userid"] = userid
	value["name"] = name
	value["username"] = username
	value["expired"] = time.Now().Add(time.Hour * 1).Format("2006-01-02 15:04:05")

	jwtKey := JWTKeys

	tokenString, _ := token.SignedString([]byte(jwtKey))

	return tokenString
}

func GetAuthorizationTokenValue(request *http.Request, param string) interface{} {
	return request.Context().Value("authorizationToken").(jwt.MapClaims)[param]
}
