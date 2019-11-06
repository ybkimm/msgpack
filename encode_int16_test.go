package msgpack

import (
	"bytes"
	"testing"
)

func TestEncoder_EncodeInt16(t *testing.T) {
	tests := []struct {
		name    string
		v       int16
		want    []byte
		wantErr bool
	}{
		{
			"a int16",
			1,
			[]byte{
				0xD1, 0x00, 0x01,
			},
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			buf := bytes.NewBuffer([]byte{})
			e := NewEncoder(buf)
			if err := e.EncodeInt16(tt.v); (err != nil) != tt.wantErr {
				t.Errorf("EncodeInt16() error = %v, wantErr %v", err, tt.wantErr)
			}
			if !bytes.Equal(buf.Bytes(), tt.want) {
				t.Errorf("EncodeInt16() got = [% X], want = [% X]", buf.Bytes(), tt.want)
			}
		})
	}
}
