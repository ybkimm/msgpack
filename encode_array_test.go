package msgpack

import (
	"bytes"
	"testing"
)

func TestEncoder_encodeArray(t *testing.T) {
	tests := []struct {
		name    string
		arr     Array
		want    []byte
		wantErr bool
	}{
		{
			"string array",
			&StringArray{
				"Hello", "World!",
			},
			[]byte{
				0x92, 0xA5, 0x48, 0x65, 0x6C, 0x6C, 0x6F, 0xA6,
				0x57, 0x6F, 0x72, 0x6C, 0x64, 0x21,
			},
			false,
		},
		{
			"int array",
			&IntArray{
				1, 2, 3,
			},
			[]byte{0x93, 0x01, 0x02, 0x03},
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := NewEncoder(nil).encodeArray(tt.arr)
			if (err != nil) != tt.wantErr {
				t.Errorf("encodeArray() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !bytes.Equal(got, tt.want) {
				t.Errorf("encodeArray() got = [% X], want = [% X]", got, tt.want)
			}
		})
	}
}
