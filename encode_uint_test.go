package msgpack

import (
	"bytes"
	"math"
	"testing"
)

func TestEncoder_encodeUint(t *testing.T) {
	tests := []struct {
		name    string
		v       uint
		want    []byte
		wantErr bool
	}{
		{
			"positive fixnum",
			0,
			[]byte{0x00},
			false,
		},
		{
			"uint8",
			math.MaxUint8,
			[]byte{0xCC, 0xFF},
			false,
		},
		{
			"uint16",
			math.MaxUint16,
			[]byte{0xCD, 0xFF, 0xFF},
			false,
		},
		{
			"uint32",
			math.MaxUint32,
			[]byte{0xCE, 0xFF, 0xFF, 0xFF, 0xFF},
			false,
		},
		{
			"uint64",
			math.MaxUint64,
			[]byte{
				0xCF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF,
				0xFF,
			},
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := NewEncoder(nil).encodeUint(tt.v)
			if (err != nil) != tt.wantErr {
				t.Errorf("encodeUint() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !bytes.Equal(got, tt.want) {
				t.Errorf("encodeUint() got = [% X], want [% X]", got, tt.want)
			}
		})
	}
}
