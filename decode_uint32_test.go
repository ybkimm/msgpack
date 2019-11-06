package msgpack

import (
	"bytes"
	"testing"
)

func TestDecoder_DecodeUint32(t *testing.T) {
	tests := []struct {
		name    string
		data    []byte
		want    uint32
		wantErr bool
	}{
		{
			"positive fixnum",
			[]byte{0x01},
			1,
			false,
		},
		{
			"uint32",
			[]byte{0xCE, 0x00, 0x00, 0x00, 0x01},
			1,
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var got uint32
			d := NewDecoder(bytes.NewReader(tt.data))
			if err := d.DecodeUint32(&got); (err != nil) != tt.wantErr {
				t.Errorf("DecodeUint32() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("DecodeUint32() got = %v, want = %v", got, tt.want)
			}
		})
	}
}
