package msgpack

import (
	"bytes"
	"testing"
)

func TestEncoder_encodeFloat32(t *testing.T) {
	tests := []struct {
		name    string
		v       float32
		want    []byte
		wantErr bool
	}{
		{
			"a float32",
			1.234,
			[]byte{0xCA, 0x3F, 0x9D, 0xF3, 0xB6},
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := NewEncoder(nil).encodeFloat32(tt.v)
			if (err != nil) != tt.wantErr {
				t.Errorf("encodeFloat32() error = %v, wantErr %v", err, tt.wantErr)
			}
			if !bytes.Equal(got, tt.want) {
				t.Errorf("encodeFloat32() got = [% X], want = [% X]", got, tt.want)
			}
		})
	}
}
