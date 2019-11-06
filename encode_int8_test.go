package msgpack

import (
	"bytes"
	"testing"
)

func TestEncoder_EncodeInt8(t *testing.T) {
	tests := []struct {
		name    string
		v       int8
		want    []byte
		wantErr bool
	}{
		{
			"a int8",
			1,
			[]byte{
				0xD0, 0x01,
			},
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			buf := bytes.NewBuffer([]byte{})
			e := NewEncoder(buf)
			if err := e.EncodeInt8(tt.v); (err != nil) != tt.wantErr {
				t.Errorf("EncodeInt8() error = %v, wantErr %v", err, tt.wantErr)
			}
			if !bytes.Equal(buf.Bytes(), tt.want) {
				t.Errorf("EncodeInt8() got = [% X], want = [% X]", buf.Bytes(), tt.want)
			}
		})
	}
}
