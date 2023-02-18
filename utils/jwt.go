package utils

import (
	"fmt"
	"github.com/golang-jwt/jwt/v4"
	"github.com/hibakun/nova-store/config"
)

func verifyJWT(t string) (*jwt.Token, error) {
	token, err := jwt.Parse(t, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(config.Config("SECRET")), nil
	})

	if err != nil {
		return nil, err
	}

	return token, nil
}

func DecodeJWT(t string) (jwt.MapClaims, error) {
	token, err := verifyJWT(t)
	if err != nil {
		return nil, err
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if token.Valid && ok {
		return claims, nil
	}

	return nil, err
}
