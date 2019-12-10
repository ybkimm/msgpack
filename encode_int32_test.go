package msgpack

import (
	"bytes"
	"testing"
)

func TestEncoder_encodeInt32(t *testing.T) {
	tests := []struct {
		name    string
		v       int32
		want    []byte
		wantErr bool
	}{
		{
			"a int32",
			1,
			[]byte{
				0xD2, 0x00, 0x00, 0x00, 0x01,
			},
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := NewEncoder(nil).encodeInt32(tt.v)
			if (err != nil) != tt.wantErr {
				t.Errorf("encodeInt32() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !bytes.Equal(got, tt.want) {
				t.Errorf("encodeInt32() got = [% X], want = [% X]", got, tt.want)
			}
		})
	}
}
