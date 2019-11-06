package msgpack

import (
	"bytes"
	"testing"
)

func TestDecoder_DecodeInt8(t *testing.T) {
	tests := []struct {
		name    string
		data    []byte
		want    int8
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
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var got int8
			d := NewDecoder(bytes.NewReader(tt.data))
			if err := d.DecodeInt8(&got); (err != nil) != tt.wantErr {
				t.Errorf("DecodeInt8() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("DecodeInt8() got = %v, want = %v", got, tt.want)
			}
		})
	}
}
