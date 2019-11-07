package msgpack

import (
	"bytes"
	"testing"
)

func TestEncoder_encodeInt8(t *testing.T) {
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
			got, err := NewEncoder(nil).encodeInt8(tt.v)
			if (err != nil) != tt.wantErr {
				t.Errorf("encodeInt8() error = %v, wantErr %v", err, tt.wantErr)
			}
			if !bytes.Equal(got, tt.want) {
				t.Errorf("encodeInt8() got = [% X], want = [% X]", got, tt.want)
			}
		})
	}
}
