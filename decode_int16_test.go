package msgpack

import (
	"bytes"
	"testing"
)

func TestDecoder_DecodeInt16(t *testing.T) {
	tests := []struct {
		name    string
		data    []byte
		want    int16
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
			"int16",
			[]byte{0xD1, 0x00, 0x01},
			1,
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var got int16
			d := NewDecoder(bytes.NewReader(tt.data))
			if err := d.DecodeInt16(&got); (err != nil) != tt.wantErr {
				t.Errorf("DecodeInt16() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("DecodeInt16() got = %v, want = %v", got, tt.want)
			}
		})
	}
}
