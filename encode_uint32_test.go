package msgpack

import (
	"bytes"
	"testing"
)

func TestEncoder_EncodeUint32(t *testing.T) {
	tests := []struct {
		name    string
		v       uint32
		want    []byte
		wantErr bool
	}{
		{
			"a uint32",
			1,
			[]byte{
				0xCE, 0x00, 0x00, 0x00, 0x01,
			},
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			buf := bytes.NewBuffer([]byte{})
			e := NewEncoder(buf)
			if err := e.EncodeUint32(tt.v); (err != nil) != tt.wantErr {
				t.Errorf("EncodeUint32() error = %v, wantErr %v", err, tt.wantErr)
			}
			if !bytes.Equal(buf.Bytes(), tt.want) {
				t.Errorf("EncodeUint32() got = [% X], want = [% X]", buf.Bytes(), tt.want)
			}
		})
	}
}
