package msgpack

import (
	"bytes"
	"testing"
)

func TestEncoder_EncodeUint64(t *testing.T) {
	tests := []struct {
		name    string
		v       uint64
		want    []byte
		wantErr bool
	}{
		{
			"a uint64",
			1,
			[]byte{
				0xCF, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x01,
			},
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			buf := bytes.NewBuffer([]byte{})
			e := NewEncoder(buf)
			if err := e.EncodeUint64(tt.v); (err != nil) != tt.wantErr {
				t.Errorf("EncodeUint64() error = %v, wantErr %v", err, tt.wantErr)
			}
			if !bytes.Equal(buf.Bytes(), tt.want) {
				t.Errorf("EncodeUint64() got = [% X], want = [% X]", buf.Bytes(), tt.want)
			}
		})
	}
}
