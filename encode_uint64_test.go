package msgpack

import (
	"bytes"
	"testing"
)

func TestEncoder_encodeUint64(t *testing.T) {
	tests := []struct {
		name    string
		v       uint64
		want    []byte
		wantErr bool
	}{
		{
			"a uint64",
			1,
			[]byte{
				0xCF, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x01,
			},
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := NewEncoder(nil).encodeUint64(tt.v)
			if (err != nil) != tt.wantErr {
				t.Errorf("encodeUint64() error = %v, wantErr %v", err, tt.wantErr)
			}
			if !bytes.Equal(got, tt.want) {
				t.Errorf("encodeUint64() got = [% X], want = [% X]", got, tt.want)
			}
		})
	}
}
