package controller

import (
	"fmt"
	_const "jwt-practice/const"
	"time"

	"github.com/golang-jwt/jwt"
)

//const secretKey = "DZrMq5mXf7KnWIR4rUY7ExrafUKOIU6U"

func GenerateJWT(email string, role string) (string, error) {

	var mySigningKey = []byte(_const.SecretKey)
	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)
	claims["authorized"] = true
	claims["email"] = email
	claims["authorized"] = role
	claims["exp"] = time.Now().Add(time.Minute * 30).Unix()
	tokenString, err := token.SignedString(mySigningKey)

	if err != nil {
		fmt.Errorf("Something went wrong", "%s", err.Error())
		return "", err
	}
	return tokenString, nil
}
