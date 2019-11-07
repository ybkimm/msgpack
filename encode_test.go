package msgpack

import (
	"encoding/json"
	"testing"
)

func BenchmarkEncoder_encodeMap(b *testing.B) {
	for i := 0; i < b.N; i++ {
		NewEncoder(nil).encodeMap(SmallStruct)
	}
}

func BenchmarkJSON_marshalMap(b *testing.B) {
	for i := 0; i < b.N; i++ {
		json.Marshal(SmallStruct)
	}
}

func BenchmarkEncoder_encodeArray(b *testing.B) {
	for i := 0; i < b.N; i++ {
		NewEncoder(nil).encodeArray(SmallArray)
	}
}

func BenchmarkJSON_marshalArray(b *testing.B) {
	for i := 0; i < b.N; i++ {
		json.Marshal(SmallArray)
	}
}
