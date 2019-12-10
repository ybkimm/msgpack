package msgpack

import (
	"bytes"
	"testing"
)

func TestEncoder_encodeUint16(t *testing.T) {
	tests := []struct {
		name    string
		v       uint16
		want    []byte
		wantErr bool
	}{
		{
			"a uint16",
			1,
			[]byte{0xCD, 0x00, 0x01},
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := NewEncoder(nil).encodeUint16(tt.v)
			if (err != nil) != tt.wantErr {
				t.Errorf("encodeUint16() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !bytes.Equal(got, tt.want) {
				t.Errorf("encodeUint16() got = [% X], want = [% X]", got, tt.want)
			}
		})
	}
}
