package msgpack

import (
	"bytes"
	"testing"
)

func TestEncoder_EncodeUint8(t *testing.T) {
	tests := []struct {
		name    string
		v       uint8
		want    []byte
		wantErr bool
	}{
		{
			"a uint8",
			1,
			[]byte{
				0xCC, 0x01,
			},
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			buf := bytes.NewBuffer([]byte{})
			e := NewEncoder(buf)
			if err := e.EncodeUint8(tt.v); (err != nil) != tt.wantErr {
				t.Errorf("EncodeUint8() error = %v, wantErr %v", err, tt.wantErr)
			}
			if !bytes.Equal(buf.Bytes(), tt.want) {
				t.Errorf("EncodeUint8() got = [% X], want = [% X]", buf.Bytes(), tt.want)
			}
		})
	}
}
