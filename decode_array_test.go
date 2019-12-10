package msgpack

import (
	"bytes"
	"reflect"
	"testing"
)

func TestDecoder_DecodeArray(t *testing.T) {
	tests := []struct {
		name    string
		orig    Array
		data    []byte
		want    Array
		wantErr bool
	}{
		{
			"string array",
			&StringArray{},
			[]byte{
				0x92, 0xA5, 0x48, 0x65, 0x6C, 0x6C, 0x6F, 0xA6,
				0x57, 0x6F, 0x72, 0x6C, 0x64, 0x21,
			},
			&StringArray{
				"Hello", "World!",
			},
			false,
		},
		{
			"int array",
			&IntArray{},
			[]byte{0x93, 0x01, 0x02, 0x03},
			&IntArray{
				1, 2, 3,
			},
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			d := NewDecoder(bytes.NewReader(tt.data))
			if err := d.DecodeArray(tt.orig); (err != nil) != tt.wantErr {
				t.Errorf("DecodeArray() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(tt.orig, tt.want) {
				t.Errorf("DecodeArray() got = %v, want = %v", tt.orig, tt.want)
			}
		})
	}
}
