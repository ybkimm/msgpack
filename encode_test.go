package msgpack

import (
	"encoding/json"
	"testing"
)

func BenchmarkEncoder_encodeMap(b *testing.B) {
	for i := 0; i < b.N; i++ {
		NewEncoder(nil).encodeMap(TestStructInstance)
	}
}

func BenchmarkJSON_marshalMap(b *testing.B) {
	for i := 0; i < b.N; i++ {
		json.Marshal(TestStructInstance)
	}
}

func BenchmarkEncoder_encodeArray(b *testing.B) {
	for i := 0; i < b.N; i++ {
		NewEncoder(nil).encodeArray(TestArrayInstance)
	}
}

func BenchmarkJSON_marshalArray(b *testing.B) {
	for i := 0; i < b.N; i++ {
		json.Marshal(TestArrayInstance)
	}
}

func BenchmarkEncoder_encodeJSON(b *testing.B) {
	for i := 0; i < b.N; i++ {
		NewEncoder(nil).encodeJSON(TestArrayDataJSON)
	}
}
