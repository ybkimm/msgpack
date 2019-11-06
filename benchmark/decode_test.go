package benchmark

import (
	"encoding/json"
	"testing"

	"github.com/francoispqt/gojay"
	msgpack2 "github.com/vmihailenco/msgpack"
	"github.com/ybkimm/msgpack"
)

func BenchmarkUnmarshalSmallData(b *testing.B) {
	for i := 0; i < b.N; i++ {
		v := &benchSmallStruct{}
		msgpack.Unmarshal(SmallData, v)
	}
}

func BenchmarkVmihailencoMsgpackUnmarshalSmallData(b *testing.B) {
	for i := 0; i < b.N; i++ {
		v := &benchSmallStruct{}
		msgpack2.Unmarshal(SmallData, v)
	}
}

func BenchmarkJSONUnmarshalSmallData(b *testing.B) {
	for i := 0; i < b.N; i++ {
		v := &benchSmallStruct{}
		json.Unmarshal(SmallDataJSON, v)
	}
}

func BenchmarkGojayUnmarshalSmallData(b *testing.B) {
	for i := 0; i < b.N; i++ {
		v := &benchSmallStruct{}
		gojay.Unmarshal(SmallDataJSON, v)
	}
}
