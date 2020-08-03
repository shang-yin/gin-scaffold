package jwt

import (
	"errors"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/spf13/viper"
)

var (
	TokenExpired     = errors.New("token is expired")
	TokenNotValidYet = errors.New("token not active yet")
	TokenMalformed   = errors.New("that's not even a token")
	TokenInvalid     = errors.New("couldn't handle this token")
)

// ParserToken
func ParserToken(tokenString string) (*CustomClaims, error) {
	signingKey := viper.GetString("jwt.signingKey")

	return at(time.Unix(0, 0), func() (*CustomClaims, error) {
		token, err := jwt.ParseWithClaims(tokenString, &CustomClaims{}, func(token *jwt.Token) (interface{}, error) {
			return []byte(signingKey), nil
		})
		if err != nil {
			if ve, ok := err.(*jwt.ValidationError); ok {
				if ve.Errors&jwt.ValidationErrorMalformed != 0 {
					return nil, TokenMalformed
				} else if ve.Errors&jwt.ValidationErrorExpired != 0 {
					return nil, TokenExpired
				} else if ve.Errors&jwt.ValidationErrorNotValidYet != 0 {
					return nil, TokenNotValidYet
				} else {
					return nil, TokenInvalid
				}
			}
		}
		if token != nil {
			if claims, ok := token.Claims.(*CustomClaims); ok && token.Valid {
				return claims, nil
			}
		}
		return nil, TokenInvalid
	})
}

func at(t time.Time, f func() (*CustomClaims, error)) (custom *CustomClaims, err error) {
	jwt.TimeFunc = func() time.Time {
		return t
	}
	custom, err = f()
	jwt.TimeFunc = time.Now
	return
}
