package msgpack

import (
	"bytes"
	"testing"
)

func TestEncoder_encodeUint8(t *testing.T) {
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
			got, err := NewEncoder(nil).encodeUint8(tt.v)
			if (err != nil) != tt.wantErr {
				t.Errorf("encodeUint8() error = %v, wantErr %v", err, tt.wantErr)
			}
			if !bytes.Equal(got, tt.want) {
				t.Errorf("encodeUint8() got = [% X], want = [% X]", got, tt.want)
			}
		})
	}
}
