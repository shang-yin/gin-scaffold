package jwt

import (
	"bytes"
	"fmt"
	"strconv"
	"testing"

	"github.com/spf13/viper"
)

var tomlConf = []byte(`
[jwt]
signingKey = "AllYourBase"  # 不填写 默认为 gin-scaffold
expiresAt = 1  # token过期时间，默认秒
`)

func TestToken(t *testing.T) {
	viper.SetConfigType("toml")
	_ = viper.ReadConfig(bytes.NewBuffer(tomlConf))
	token := GenerateToken("12")
	t.Log(token)
}

func TestParserToken(t *testing.T) {
	viper.SetConfigType("toml")
	_ = viper.ReadConfig(bytes.NewBuffer(tomlConf))
	tokenStr := "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpZCI6IjEyIiwiZXhwIjoxNTk2MDkxNzY5LCJuYmYiOjE1OTYwOTE3NTh9.mY1klWqnzDoSSBHrGv0o4ujE6Cx7bC24jF9IYVRkAvs"
	if token, err := ParserToken(tokenStr); err == nil {
		fmt.Println(token.ID, err)
	} else {
		fmt.Println(err)
	}
}

// BenchmarkGenerateToken-12    	  151341	      6947 ns/op
func BenchmarkGenerateToken(b *testing.B) {
	viper.SetConfigType("toml")
	_ = viper.ReadConfig(bytes.NewBuffer(tomlConf))

	for i := 0; i < b.N; i++ {
		GenerateToken(strconv.Itoa(i))
	}
}

func BenchmarkParserToken(b *testing.B) {
	viper.SetConfigType("toml")
	_ = viper.ReadConfig(bytes.NewBuffer(tomlConf))
	tokenStr := "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpZCI6IjEyIiwiZXhwIjoxNTAwMH0.CY7iOyRqf5-drskIxzR04zxQwoqXIXU6fbQ77pHupoU"
	for i := 0; i < b.N; i++ {
		_, _ = ParserToken(tokenStr)
	}
}
