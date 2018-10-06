package model

import (
	"myproject/common"

	jwt "github.com/dgrijalva/jwt-go"
)

type JwtToken struct {
	Token string `json:"token"`
}

func GetToken(user common.User) (JwtToken, error) {

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"username": user.Username,
		"password": user.Password,
	})
	tokenString, error := token.SignedString([]byte("secret"))
	if error != nil {
		return JwtToken{}, error
	}
	return JwtToken{Token: tokenString}, nil
}
