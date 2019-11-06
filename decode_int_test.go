package msgpack

import (
	"bytes"
	"testing"
)

func TestDecoder_DecodeInt(t *testing.T) {
	tests := []struct {
		name    string
		data    []byte
		want    int
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
			"int8",
			[]byte{0xD0, 0x01},
			1,
			false,
		},
		{
			"int16",
			[]byte{0xD1, 0x00, 0x01},
			1,
			false,
		},
		{
			"int32",
			[]byte{0xD2, 0x00, 0x00, 0x00, 0x01},
			1,
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
			var got int
			d := NewDecoder(bytes.NewReader(tt.data))
			if err := d.DecodeInt(&got); (err != nil) != tt.wantErr {
				t.Errorf("DecodeInt() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("DecodeInt() got = %v, want = %v", got, tt.want)
			}
		})
	}
}
