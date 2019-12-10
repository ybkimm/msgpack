package msgpack

import (
	"bytes"
	"testing"
)

func TestEncoder_encodeBinary(t *testing.T) {
	tests := []struct {
		name    string
		p       []byte
		want    []byte
		wantErr bool
	}{
		{
			"a binary",
			[]byte{0x11, 0x22, 0x33, 0x44},
			[]byte{0xC4, 0x04, 0x11, 0x22, 0x33, 0x44},
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := NewEncoder(nil).encodeBinary(tt.p)
			if (err != nil) != tt.wantErr {
				t.Errorf("encodeBinary() error = %v, wantErr %v", err, tt.wantErr)
			}
			if !bytes.Equal(got, tt.want) {
				t.Errorf("encodeBinary() got = [% X], want [% X]", got, tt.want)
			}
		})
	}
}
