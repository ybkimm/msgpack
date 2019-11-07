package msgpack

import (
	"bytes"
	"testing"
)

var benchDataDecodeBool = []byte{0xC3}

func TestDecoder_DecodeBool(t *testing.T) {
	tests := []struct {
		name    string
		data    []byte
		want    bool
		wantErr bool
	}{
		{
			"true",
			[]byte{0xC3},
			true,
			false,
		},
		{
			"false",
			[]byte{0xC2},
			false,
			false,
		},
		{
			"else",
			[]byte{0x00},
			false,
			true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var got bool
			d := NewDecoder(bytes.NewReader(tt.data))
			if err := d.DecodeBool(&got); (err != nil) != tt.wantErr {
				t.Errorf("DecodeBool() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("DecodeBool() got = %v, want = %v", got, tt.want)
			}
		})
	}
}
