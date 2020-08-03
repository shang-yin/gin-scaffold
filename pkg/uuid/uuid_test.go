package uuid

import (
	"fmt"
	"testing"
)

func TestUUID(t *testing.T) {
	uuid := UUID()
	fmt.Println("uuid = ", uuid)
}

func BenchmarkUUID(b *testing.B) {
	for i := 0; i <= b.N; i++ {
		UUID()
	}
}
