package database

import (
	"bytes"
	"log"
	"sync"
	"testing"

	"github.com/spf13/viper"
)

var tomlConf = []byte(`
[database]
driver = "mysql"  # 不传 默认为MYSQL
server = "127.0.0.1"
port = 3306
username = "root"
password = "123456"
dbname = "gin-scaffold"
prefix = "g_"  # 表前缀
charset = "utf8mb4"  # 字符集设置
# 一般建议maxIdleConn的值为MaxOpenConn的1/2
maxIdleConn = 50  # 空闲连接池中连接的最大数量
maxOpenConn = 100 # 打开数据库连接的最大数量
maxLifeTime = 600 # 连接可复用的最大时间 默认单位 秒
`)

// 链接池测试
func TestInitDB(t *testing.T) {
	viper.SetConfigType("toml")
	_ = viper.ReadConfig(bytes.NewBuffer(tomlConf))

	InitDB()
	var sy sync.WaitGroup
	// 开启20个goroutine
	for i := 0; i < 20; i++ {
		sy.Add(1)
		go func(i int) {
			oneWorker(i)
			sy.Done()
		}(i)
	}
	sy.Wait()

}
func oneWorker(i int) {

	var connectionId int
	row := DB.Raw("select CONNECTION_ID()").Row()
	err := row.Scan(&connectionId)
	if err != nil {
		log.Println("query connection id failed:", err)
		return
	}

	log.Println("worker:", i, ", connection id:", connectionId)

	var result int
	row = DB.Raw("select sleep(10)").Row()
	err = row.Scan(&result)
	if err != nil {
		log.Println("query sleep connection id faild:", err)
		return
	}
}
