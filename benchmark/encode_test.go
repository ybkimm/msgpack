package benchmark

import (
	"encoding/json"
	"testing"

	"github.com/ybkimm/msgpack"
)

func BenchmarkJSONMarshalSmallData(b *testing.B) {
	for i := 0; i < b.N; i++ {
		json.Marshal(SmallData)
	}
}

func BenchmarkMarshalSmallData(b *testing.B) {
	for i := 0; i < b.N; i++ {
		msgpack.Marshal(SmallData)
	}
}
