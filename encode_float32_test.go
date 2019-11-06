package msgpack

import (
	"bytes"
	"testing"
)

func TestEncoder_EncodeFloat32(t *testing.T) {
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
			buf := bytes.NewBuffer([]byte{})
			e := NewEncoder(buf)
			if err := e.EncodeFloat32(tt.v); (err != nil) != tt.wantErr {
				t.Errorf("EncodeFloat32() error = %v, wantErr %v", err, tt.wantErr)
			}
			if !bytes.Equal(buf.Bytes(), tt.want) {
				t.Errorf("EncodeFloat32() got = [% X], want = [% X]", buf.Bytes(), tt.want)
			}
		})
	}
}
