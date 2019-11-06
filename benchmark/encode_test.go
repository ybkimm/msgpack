package benchmark

import (
	"encoding/json"
	"testing"

	"github.com/francoispqt/gojay"
	msgpack2 "github.com/vmihailenco/msgpack"
	"github.com/ybkimm/msgpack"
)

func BenchmarkMarshalSmallData(b *testing.B) {
	for i := 0; i < b.N; i++ {
		msgpack.Marshal(SmallStruct)
	}
}

func BenchmarkVmihailencoMsgpackMarshalSmallData(b *testing.B) {
	for i := 0; i < b.N; i++ {
		msgpack2.Marshal(SmallStruct)
	}
}

func BenchmarkJSONMarshalSmallData(b *testing.B) {
	for i := 0; i < b.N; i++ {
		json.Marshal(SmallStruct)
	}
}

func BenchmarkGojayMarshalSmallData(b *testing.B) {
	for i := 0; i < b.N; i++ {
		gojay.Marshal(SmallStruct)
	}
}
