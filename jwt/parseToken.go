package jwt

import (
	"errors"
	"github.com/golang-jwt/jwt/v4"
)

func UnSecret() jwt.Keyfunc {
	return func(token *jwt.Token) (interface{}, error) {
		return []byte(secret), nil
	}
}

func ParseToken(tokens string) (*UserClaims, error) {
	token, err := jwt.ParseWithClaims(tokens, &UserClaims{}, UnSecret())
	if err != nil {
		if ve, ok := err.(*jwt.ValidationError); ok {
			if ve.Errors&jwt.ValidationErrorMalformed != 0 {
				return nil, errors.New("that's not even a token")
			} else if ve.Errors&jwt.ValidationErrorExpired != 0 {
				return nil, errors.New("token is expired")
			} else if ve.Errors&jwt.ValidationErrorNotValidYet != 0 {
				return nil, errors.New("token not active yet")
			} else {
				return nil, errors.New("couldn't handle this token")
			}
		}
	}
	if claims, ok := token.Claims.(*UserClaims); ok && token.Valid {
		return claims, nil
	}
	return nil, errors.New("couldn't handle this token")
}
