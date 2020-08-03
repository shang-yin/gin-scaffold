package conf

import (
	"fmt"
	"testing"

	"github.com/spf13/viper"
)

func TestGetConf(t *testing.T) {
	viper.AddConfigPath("test_conf")
	InitConf()
	conf := GetConf()
	fmt.Println(conf)
}

func BenchmarkGetConf(b *testing.B) {
	viper.AddConfigPath("test_conf")
	InitConf()

	for i := 0; i < b.N; i++ {
		GetConf()
	}

}
