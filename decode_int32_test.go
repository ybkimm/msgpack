package msgpack

import (
	"bytes"
	"testing"
)

func TestDecoder_DecodeInt32(t *testing.T) {
	tests := []struct {
		name    string
		data    []byte
		want    int32
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
			"int32",
			[]byte{0xD2, 0x00, 0x00, 0x00, 0x01},
			1,
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var got int32
			d := NewDecoder(bytes.NewReader(tt.data))
			if err := d.DecodeInt32(&got); (err != nil) != tt.wantErr {
				t.Errorf("DecodeInt32() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("DecodeInt32() got = %v, want = %v", got, tt.want)
			}
		})
	}
}
