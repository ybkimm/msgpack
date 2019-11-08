package msgpack

import (
	"encoding/json"
	"testing"
)

func BenchmarkDecoder_decodeMap(b *testing.B) {
	var result = TestStruct{}
	for i := 0; i < b.N; i++ {
		NewBytesDecoder(TestStructData).Decode(&result)
	}
}

func BenchmarkDecoder_decodeArray(b *testing.B) {
	var result = TestArray{}
	for i := 0; i < b.N; i++ {
		NewBytesDecoder(TestArrayData).Decode(&result)
	}
}

func BenchmarkJSON_decode(b *testing.B) {
	var result = TestStruct{}
	for i := 0; i < b.N; i++ {
		json.Unmarshal(TestStructDataJSON, &result)
	}
}
