package msgpack

import (
	"bytes"
	"testing"
)

func TestDecoder_DecodeUint16(t *testing.T) {
	tests := []struct {
		name    string
		data    []byte
		want    uint16
		wantErr bool
	}{
		{
			"positive fixnum",
			[]byte{0x01},
			1,
			false,
		},
		{
			"uint16",
			[]byte{0xCD, 0x00, 0x01},
			1,
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var got uint16
			d := NewDecoder(bytes.NewReader(tt.data))
			if err := d.DecodeUint16(&got); (err != nil) != tt.wantErr {
				t.Errorf("DecodeUint16() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("DecodeUint16() got = %v, want = %v", got, tt.want)
			}
		})
	}
}
