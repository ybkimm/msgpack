package msgpack

import (
	"bytes"
	"testing"
)

func TestDecoder_DecodeInt64(t *testing.T) {
	tests := []struct {
		name    string
		data    []byte
		want    int64
		wantErr bool
	}{
		{
			"positive fixnum",
			[]byte{0x01},
			1,
			false,
		},
		{
			"negative fixnum",
			[]byte{0xFF},
			-1,
			false,
		},
		{
			"int64",
			[]byte{
				0xD3, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
				0x01,
			},
			1,
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var got int64
			d := NewDecoder(bytes.NewReader(tt.data))
			if err := d.DecodeInt64(&got); (err != nil) != tt.wantErr {
				t.Errorf("DecodeInt64() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("DecodeInt64() got = %v, want = %v", got, tt.want)
			}
		})
	}
}
