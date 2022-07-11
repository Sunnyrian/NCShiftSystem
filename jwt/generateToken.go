package jwt

import (
	"github.com/golang-jwt/jwt/v4"
	"time"
)

type UserClaims struct {
	StuID string `json:"stuID"`
	jwt.RegisteredClaims
	Admin bool
}

// secret 后期应该抽离到配置文件
var secret = []byte("学生网络管理协会")

// GenerateToken 生成 Token
func GenerateToken(ID string, admin bool) (tokenString string, err error) {
	claim := UserClaims{
		StuID: ID,
		Admin: admin,
		RegisteredClaims: jwt.RegisteredClaims{
			// 设置过期时间为 3 小时 这里其实算魔法值，应该抽离成变量
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(3 * time.Hour * time.Duration(1))),
			IssuedAt: jwt.NewNumericDate(time.Now()),
			NotBefore: jwt.NewNumericDate(time.Now()),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claim)
	tokenString, err = token.SignedString(secret)
	return tokenString, err
}
