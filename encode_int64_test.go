package msgpack

import (
	"bytes"
	"testing"
)

func TestEncoder_EncodeInt64(t *testing.T) {
	tests := []struct {
		name    string
		v       int64
		want    []byte
		wantErr bool
	}{
		{
			"a int64",
			1,
			[]byte{
				0xD3, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
				0x01,
			},
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			buf := bytes.NewBuffer([]byte{})
			e := NewEncoder(buf)
			if err := e.EncodeInt64(tt.v); (err != nil) != tt.wantErr {
				t.Errorf("EncodeInt64() error = %v, wantErr %v", err, tt.wantErr)
			}
			if !bytes.Equal(buf.Bytes(), tt.want) {
				t.Errorf("EncodeInt64() got = [% X], want = [% X]", buf.Bytes(), tt.want)
			}
		})
	}
}
