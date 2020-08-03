package jwt

import (
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/spf13/viper"
)

type CustomClaims struct {
	ID string `json:"id"`
	jwt.StandardClaims
}

// GenerateToken 生成TOKEN
func GenerateToken(id string) string {
	signingKey := viper.GetString("jwt.signingKey")
	expiresAt := viper.GetInt64("jwt.expiresAt")

	claims := CustomClaims{
		ID: id,
		StandardClaims: jwt.StandardClaims{
			NotBefore: time.Now().Unix() - 10, // 放置立即生效失败
			ExpiresAt: time.Now().Unix() + expiresAt,
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	signedString, err := token.SignedString([]byte(signingKey))
	if err != nil {
		panic(err)
	}
	return signedString
}
