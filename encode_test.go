package msgpack

import (
	"testing"
)

func BenchmarkEncoder_encode(b *testing.B) {
	for i := 0; i < b.N; i++ {
		NewEncoder(nil).encode(SmallStruct)
	}
}
