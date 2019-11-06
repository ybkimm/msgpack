package msgpack

import (
	"bytes"
	"math"
	"testing"
)

func TestDecoder_DecodeFloat32(t *testing.T) {
	tests := []struct {
		name    string
		data    []byte
		want    float32
		wantErr bool
	}{
		{
			"float32",
			[]byte{0xCA, 0x3F, 0x9D, 0xF3, 0xB6},
			1.234,
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var got float32
			d := NewDecoder(bytes.NewReader(tt.data))
			if err := d.DecodeFloat32(&got); (err != nil) != tt.wantErr {
				t.Errorf("DecodeFloat32() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if math.Abs(float64(got-tt.want)) > 0.01 {
				t.Errorf("DecodeFloat32() got = %v, want = %v", got, tt.want)
			}
		})
	}
}
