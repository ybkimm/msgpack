package msgpack

import (
	"bytes"
	"testing"
)

func TestEncoder_encodeInt16(t *testing.T) {
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
			got, err := NewEncoder(nil).encodeInt16(tt.v)
			if (err != nil) != tt.wantErr {
				t.Errorf("encodeInt16() error = %v, wantErr %v", err, tt.wantErr)
			}
			if !bytes.Equal(got, tt.want) {
				t.Errorf("encodeInt16() got = [% X], want = [% X]", got, tt.want)
			}
		})
	}
}
