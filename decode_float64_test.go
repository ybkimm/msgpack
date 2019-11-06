package msgpack

import (
	"bytes"
	"math"
	"testing"
)

func TestDecoder_DecodeFloat64(t *testing.T) {
	tests := []struct {
		name    string
		data    []byte
		want    float64
		wantErr bool
	}{
		{
			"float64",
			[]byte{
				0xCB, 0x3F, 0xF3, 0xBE, 0x76, 0xC8, 0xB4, 0x39,
				0x58,
			},
			1.234,
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var got float64
			d := NewDecoder(bytes.NewReader(tt.data))
			if err := d.DecodeFloat64(&got); (err != nil) != tt.wantErr {
				t.Errorf("DecodeFloat32() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if math.Abs(got-tt.want) > 0.01 {
				t.Errorf("DecodeFloat32() got = %v, want = %v", got, tt.want)
			}
		})
	}
}
