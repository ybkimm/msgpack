package msgpack

import (
	"bytes"
	"testing"
)

func TestEncoder_EncodeInt(t *testing.T) {
	tests := []struct {
		name    string
		v       int
		want    []byte
		wantErr bool
	}{
		{
			"positive fixnum",
			0,
			[]byte{0x00},
			false,
		},
		{
			"negative fixnum",
			-1,
			[]byte{0xFF},
			false,
		},
		{
			"int8",
			-128,
			[]byte{0xD0, 0x80},
			false,
		},
		{
			"int16",
			-32768,
			[]byte{0xD1, 0x80, 0x00},
			false,
		},
		{
			"int32",
			-2147483648,
			[]byte{0xD2, 0x80, 0x00, 0x00, 0x00},
			false,
		},
		{
			"int64",
			-9223372036854775808,
			[]byte{0xD3, 0x80, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00},
			false,
		},
		{
			"uint8",
			1<<8 - 1,
			[]byte{0xCC, 0xFF},
			false,
		},
		{
			"uint16",
			1 << 8,
			[]byte{0xCD, 0x01, 0x00},
			false,
		},
		{
			"uint32",
			1 << 16,
			[]byte{0xCE, 0x00, 0x01, 0x00, 0x00},
			false,
		},
		{
			"uint64",
			1 << 32,
			[]byte{0xCF, 0x00, 0x00, 0x00, 0x01, 0x00, 0x00, 0x00, 0x00},
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			buf := bytes.NewBuffer(make([]byte, 0, 512))
			e := NewEncoder(buf)
			if err := e.EncodeInt(tt.v); (err != nil) != tt.wantErr {
				t.Errorf("EncodeInt() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !bytes.Equal(buf.Bytes(), tt.want) {
				t.Errorf("EncodeInt() got = [% X], want [% X]", buf.Bytes(), tt.want)
			}
		})
	}
}
