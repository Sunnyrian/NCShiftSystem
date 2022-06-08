package jwt

import (
	"errors"
	"github.com/golang-jwt/jwt/v4"
	"time"
)

func RefreshToken(tokenStr string) (string, error) {
	jwt.TimeFunc = func() time.Time {
		return time.Unix(0,0)
	}

	token, err := jwt.ParseWithClaims(tokenStr, &UserClaims{}, func(token *jwt.Token) (interface{}, error) {
		return secret, nil
	})
	if err != nil {
		return "", err
	}
	if claims, ok := token.Claims.(*UserClaims); ok && token.Valid{
		jwt.TimeFunc = time.Now
		claims.RegisteredClaims.ExpiresAt = jwt.NewNumericDate(time.Now().Add(time.Hour * 3))
		return GenerateToken(claims.StuID)
	}
	return "", errors.New("couldn't handle this token")
}