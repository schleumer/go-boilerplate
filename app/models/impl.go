package models

import (
	"github.com/dgrijalva/jwt-go"
	"../system"
)

func (u *User) SignToken() string {
	// Create a new token object, specifying signing method and the claims
	// you would like it to contain.
	token := jwt.NewWithClaims(jwt.SigningMethodHS512, jwt.MapClaims{
		"sub": u.ID,
	})

	tokenString, _ := token.SignedString(system.JwtSecret)

	return tokenString
}
