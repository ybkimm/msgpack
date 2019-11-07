package msgpack

import (
	"bytes"
	"testing"
)

func TestEncoder_encodeUint32(t *testing.T) {
	tests := []struct {
		name    string
		v       uint32
		want    []byte
		wantErr bool
	}{
		{
			"a uint32",
			1,
			[]byte{0xCE, 0x00, 0x00, 0x00, 0x01},
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := NewEncoder(nil).encodeUint32(tt.v)
			if (err != nil) != tt.wantErr {
				t.Errorf("encodeUint32() error = %v, wantErr %v", err, tt.wantErr)
			}
			if !bytes.Equal(got, tt.want) {
				t.Errorf("encodeUint32() got = [% X], want = [% X]", got, tt.want)
			}
		})
	}
}
