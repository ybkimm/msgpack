package msgpack

import (
	"bytes"
	"testing"
)

func TestDecoder_DecodeUint8(t *testing.T) {
	tests := []struct {
		name    string
		data    []byte
		want    uint8
		wantErr bool
	}{
		{
			"positive fixnum",
			[]byte{0x01},
			1,
			false,
		},
		{
			"uint8",
			[]byte{0xCC, 0x01},
			1,
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var got uint8
			d := NewDecoder(bytes.NewReader(tt.data))
			if err := d.DecodeUint8(&got); (err != nil) != tt.wantErr {
				t.Errorf("DecodeUint8() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("DecodeUint8() got = %v, want = %v", got, tt.want)
			}
		})
	}
}
