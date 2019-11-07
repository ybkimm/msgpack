package msgpack

import (
	"bytes"
	"testing"
)

func TestEncoder_encodeInt64(t *testing.T) {
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
			got, err := NewEncoder(nil).encodeInt64(tt.v)
			if (err != nil) != tt.wantErr {
				t.Errorf("encodeInt64() error = %v, wantErr %v", err, tt.wantErr)
			}
			if !bytes.Equal(got, tt.want) {
				t.Errorf("encodeInt64() got = [% X], want = [% X]", got, tt.want)
			}
		})
	}
}
