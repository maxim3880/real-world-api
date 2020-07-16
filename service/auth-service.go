package service

import (
	"time"

	"github.com/dgrijalva/jwt-go"
)

const signatureKey = "application_signature_key"

//GenerateJwtTocken return correct auth jwt tocken by user name
func GenerateJwtTocken(username string, id int) string {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user":   username,
		"exp":    time.Now().Add(time.Hour * time.Duration(1)).Unix(),
		"iat":    time.Now().Unix(),
		"userID": id,
	})
	tokenString, err := token.SignedString(getKeyAsByteArray())
	if err == nil {
		return tokenString
	}
	return ""
}

//GetJwtValidationKey return signature key
func GetJwtValidationKey(token *jwt.Token) (interface{}, error) {
	return getKeyAsByteArray(), nil
}

func getKeyAsByteArray() []byte {
	return []byte(signatureKey)
}
