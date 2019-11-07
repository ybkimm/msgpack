package msgpack

import (
	"bytes"
	"encoding/json"
	"testing"
)

func BenchmarkDecoder_decodeMap(b *testing.B) {
	var result = TestStruct{}
	for i := 0; i < b.N; i++ {
		NewDecoder(bytes.NewReader(TestStructData)).Decode(&result)
	}
}

func BenchmarkDecoder_decodeArray(b *testing.B) {
	var result = TestArray{}
	for i := 0; i < b.N; i++ {
		NewDecoder(bytes.NewReader(TestArrayData)).Decode(&result)
	}
}

func BenchmarkJSON_decode(b *testing.B) {
	var result = TestStruct{}
	for i := 0; i < b.N; i++ {
		json.Unmarshal(TestStructDataJSON, &result)
	}
}
