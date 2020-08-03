package logging

import (
	"fmt"
	"time"

	rotatelogs "github.com/lestrrat-go/file-rotatelogs"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

var (
	log      = logrus.New()
	logEntry *logrus.Entry
)

// InitLog .
func init() {
	fmt.Println("loading logging .....")
	logPath := viper.GetString("log.path")
	/* 日志轮转相关函数
	`WithLinkName` 为最新的日志建立软连接
	`WithRotationTime` 设置日志分割的时间，隔多久分割一次
	WithMaxAge 和 WithRotationCount二者只能设置一个
	  `WithMaxAge` 设置文件清理前的最长保存时间
	  `WithRotationCount` 设置文件清理前最多保存的个数
	*/
	rl, err := rotatelogs.New(
		fmt.Sprintf(logPath, "%Y%m%d"),
		rotatelogs.WithMaxAge(time.Duration(viper.GetInt64("log.WithMaxAge"))*time.Hour),
		rotatelogs.WithRotationTime(time.Duration(viper.GetInt64("log.RotationTime"))*time.Hour),
	)
	if err != nil {
		panic(err)
	}
	log.SetOutput(rl)
	log.SetFormatter(&logrus.TextFormatter{
		TimestampFormat:  "2006-01-02 15:04:05",
		DisableTimestamp: false,
		FieldMap:         nil,
	})
	logEntry = log.WithFields(logrus.Fields{
		"topic": viper.GetString("title"),
	})
}

// Info .
func Info(msg interface{}) {
	logEntry.Info(msg)
}

// Info .
func Debug(msg interface{}) {
	logEntry.Debug(msg)
}

// Error
func Error(msg interface{}) {
	logEntry.Error(msg)
}
