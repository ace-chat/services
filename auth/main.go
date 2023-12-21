package auth

import (
	"ace/model"
	"fmt"
	"github.com/golang-jwt/jwt/v5"
	"strconv"
	"time"
)

type Claims struct {
	jwt.RegisteredClaims
	model.User
}

var sign = []byte("ace")

func GenerateToken(id uint, user model.User) (string, error) {
	issuer := "ace"
	claims := Claims{
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour * 24)),
			Issuer:    issuer,
			ID:        strconv.Itoa(int(id)),
		},
		User: user,
	}

	token, err := jwt.NewWithClaims(jwt.SigningMethodHS256, claims).SignedString(sign)
	return token, err
}

func ParseToken(tokenStr string) (*Claims, error) {
	token, err := jwt.ParseWithClaims(tokenStr, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		return sign, nil
	})
	if err != nil {
		return nil, err
	}

	if claims, ok := token.Claims.(*Claims); ok && token.Valid {
		return claims, nil
	} else {
		return nil, fmt.Errorf("valid error")
	}
}
