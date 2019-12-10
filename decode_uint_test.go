package msgpack

import (
	"bytes"
	"testing"
)

func TestDecoder_DecodeUint(t *testing.T) {
	tests := []struct {
		name    string
		data    []byte
		want    uint
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
		{
			"uint16",
			[]byte{0xCD, 0x00, 0x01},
			1,
			false,
		},
		{
			"uint32",
			[]byte{0xCE, 0x00, 0x00, 0x00, 0x01},
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
			var got uint
			d := NewDecoder(bytes.NewReader(tt.data))
			if err := d.DecodeUint(&got); (err != nil) != tt.wantErr {
				t.Errorf("DecodeUint() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("DecodeUint() got = %v, want = %v", got, tt.want)
			}
		})
	}
}
