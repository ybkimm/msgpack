package msgpack

import (
	"bytes"
	"testing"
)

func TestDecoder_DecodeUint64(t *testing.T) {
	tests := []struct {
		name    string
		data    []byte
		want    uint64
		wantErr bool
	}{
		{
			"positive fixnum",
			[]byte{0x01},
			1,
			false,
		},
		{
			"uint64",
			[]byte{
				0xCF, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
				0x01,
			},
			1,
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var got uint64
			d := NewDecoder(bytes.NewReader(tt.data))
			if err := d.DecodeUint64(&got); (err != nil) != tt.wantErr {
				t.Errorf("DecodeUint64() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("DecodeUint64() got = %v, want = %v", got, tt.want)
			}
		})
	}
}
