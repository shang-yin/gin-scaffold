package nosql

import (
	"fmt"
	"gin-scaffold/pkg/conf"

	"github.com/go-redis/redis/v8"
)

var (
	// Db .
	Db *redis.Client
)

func init() {
	if ok := conf.Config.IsSet("redis"); !ok {
		panic("please set redis addr parameters first")
	}
	redisConf := conf.Config.GetStringMapString("redis")
	fmt.Println("Addr = ", redisConf["addr"])
	Db = redis.NewClient(&redis.Options{
		Addr:     redisConf["addr"],
		Password: redisConf["password"],
		DB:       0,
	})
}
