package middleware

import (
	"gin-scaffold/pkg/comm"
	"gin-scaffold/pkg/e"
	"gin-scaffold/pkg/jwt"

	"github.com/gin-gonic/gin"
)

// Auth 权限校验
func Auth() gin.HandlerFunc {
	return func(c *gin.Context) {
		authorization := c.Request.Header.Get("Authorization")
		if authorization == "" {
			comm.ReturnJSON(c, e.ErrTokenIsNull, nil)
			c.Abort()
			return
		}
		token, err := jwt.ParserToken(authorization)
		if err != nil {
			if err == jwt.TokenInvalid || err == jwt.TokenNotValidYet || err == jwt.TokenMalformed {
				comm.ReturnJSON(c, e.ErrInvalidAuth, nil)
			} else if err == jwt.TokenExpired {
				comm.ReturnJSON(c, e.ErrTokenExpired)
			}
			c.Abort()
			return
		}
		if token != nil {
			c.Set("id", token.ID)
		}
		c.Next()
	}
}
