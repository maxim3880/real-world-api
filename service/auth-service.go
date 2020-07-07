package service

import (
	"time"

	"github.com/dgrijalva/jwt-go"
)

const signatureKey = "application_signature_key"

//GenerateJwtTocken return correct auth jwt tocken by user name
func GenerateJwtTocken(username string) string {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user": username,
		"exp":  time.Now().Add(time.Hour * time.Duration(1)).Unix(),
		"iat":  time.Now().Unix(),
	})
	tokenString, err := token.SignedString([]byte(signatureKey))
	if err == nil {
		return tokenString
	}
	return ""
}

//ParseJwtAuthHeader return claims from jwt tocken
func ParseJwtAuthHeader(authHeader string) {

}
