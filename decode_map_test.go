package msgpack

import (
	"bytes"
	"reflect"
	"testing"
)

type decodeteststruct struct {
	foo int
	bar string
}

func (t *decodeteststruct) MarshalMsgpackMap(e *Encoder) {
	panic("unimplemented")
}

func (t *decodeteststruct) UnmarshalMsgpackMap(d *Decoder, key string) error {
	switch key {
	case "foo":
		return d.DecodeInt(&t.foo)

	case "bar":
		return d.DecodeString(&t.bar)
	}

	return &ErrUnknownKey{key}
}

func (t *decodeteststruct) KeySize() uint32 {
	return 2
}

func TestDecoder_DecodeMap(t *testing.T) {
	tests := []struct {
		name    string
		orig    Map
		data    []byte
		want    Map
		wantErr bool
	}{
		{
			"object",
			&decodeteststruct{},
			[]byte{
				0x82, 0xA3, 0x66, 0x6F, 0x6F, 0x0A, 0xA3, 0x62,
				0x61, 0x72, 0xAD, 0x48, 0x65, 0x6C, 0x6C, 0x6F,
				0x2C, 0x20, 0x57, 0x6F, 0x72, 0x6C, 0x64, 0x21,
			},
			&decodeteststruct{
				foo: 10,
				bar: "Hello, World!",
			},
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			d := NewDecoder(bytes.NewReader(tt.data))
			if err := d.DecodeMap(tt.orig); (err != nil) != tt.wantErr {
				t.Errorf("DecodeMap() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(tt.orig, tt.want) {
				t.Errorf("DecodeMap() got = %v, want %v", tt.orig, tt.want)
			}
		})
	}
}
