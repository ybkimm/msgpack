package msgpack

import (
	"bytes"
	"testing"
)

func TestEncoder_EncodeBinary(t *testing.T) {
	tests := []struct {
		name    string
		p       []byte
		want    []byte
		wantErr bool
	}{
		{
			"a binary",
			[]byte{0x11, 0x22, 0x33, 0x44},
			[]byte{0xC4, 0x04, 0x11, 0x22, 0x33, 0x44},
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			buf := bytes.NewBuffer(make([]byte, 0, 512))
			e := NewEncoder(buf)
			if err := e.EncodeBinary(tt.p); (err != nil) != tt.wantErr {
				t.Errorf("EncodeBinary() error = %v, wantErr %v", err, tt.wantErr)
			}
			if !bytes.Equal(buf.Bytes(), tt.want) {
				t.Errorf("EncodeBinary() got = [% X], want [% X]", buf.Bytes(), tt.want)
			}
		})
	}
}
