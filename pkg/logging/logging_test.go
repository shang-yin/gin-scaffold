package logging

import (
	"fmt"
	"testing"

	"github.com/spf13/viper"
)

type testLog struct {
	Name string `json:"name"`
}

func TestInfo(t *testing.T) {
	viper.AddConfigPath("test_conf")
	Info("test logrus")
	Info(map[string]string{"aaa": "222"})
	Info(fmt.Sprintf("%+v", testLog{Name: "test"}))
}

func BenchmarkInfo(b *testing.B) {
	for i := 0; i < b.N; i++ {
		// logEntry BenchmarkInfo-12    	  169213	      6671 ns/op
		// logHook  BenchmarkInfo-12    	  161727	      7563 ns/op
		Info("test logrus")
	}
}
