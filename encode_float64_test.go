package msgpack

import (
	"bytes"
	"testing"
)

func TestEncoder_encodeFloat64(t *testing.T) {
	tests := []struct {
		name    string
		v       float64
		want    []byte
		wantErr bool
	}{
		{
			"a float64",
			1.234,
			[]byte{
				0xCB, 0x3F, 0xF3, 0xBE, 0x76, 0xC8, 0xB4, 0x39,
				0x58,
			},
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := NewEncoder(nil).encodeFloat64(tt.v)
			if (err != nil) != tt.wantErr {
				t.Errorf("encodeFloat64() error = %v, wantErr %v", err, tt.wantErr)
			}
			if !bytes.Equal(got, tt.want) {
				t.Errorf("encodeFloat64() got = [% X], want = [% X]", got, tt.want)
			}
		})
	}
}
