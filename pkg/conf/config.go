package conf

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

var Config = viper.GetViper()

func init() {
	fmt.Println("loading config .....")
	viper.SetConfigName(gin.Mode())
	viper.SetConfigType("toml")
	viper.AddConfigPath("conf")
	if err := viper.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			// Config file not found; ignore error if desired
			fmt.Println("配置文件未找到")
		} else {
			// Config file was found but another error was produced

			fmt.Println("another err = ", err)
		}
	}
}

func GetConf() map[string]string {
	return viper.GetStringMapString("database")
}
