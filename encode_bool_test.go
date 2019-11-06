package msgpack

import (
	"bytes"
	"testing"
)

func TestEncoder_EncodeBool(t *testing.T) {
	tests := []struct {
		name    string
		v       bool
		want    []byte
		wantErr bool
	}{
		{
			"true",
			true,
			[]byte{0xC3},
			false,
		},
		{
			"false",
			false,
			[]byte{0xC2},
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			buf := bytes.NewBuffer(make([]byte, 0, 512))
			e := NewEncoder(buf)
			if err := e.EncodeBool(tt.v); (err != nil) != tt.wantErr {
				t.Errorf("EncodeBool() error = %v, wantErr %v", err, tt.wantErr)
			}
			if !bytes.Equal(buf.Bytes(), tt.want) {
				t.Errorf("EncodeBool() got = [% X], want [% X]", buf.Bytes(), tt.want)
			}
		})
	}
}
