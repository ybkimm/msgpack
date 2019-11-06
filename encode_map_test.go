package msgpack

import (
	"bytes"
	"testing"
)

var _ Map = (*encodeteststruct)(nil)

type encodeteststruct struct {
	foo int
	bar string
}

func (t *encodeteststruct) MarshalMsgpackMap(e *Encoder) {
	e.PutIntKey("foo", t.foo)
	e.PutStringKey("bar", t.bar)
}

func (t *encodeteststruct) UnmarshalMsgpackMap(e *Decoder, key string) error {
	panic("unimplemented")
}

func (t *encodeteststruct) KeySize() uint32 {
	return 2
}

func TestEncoder_EncodeMap(t *testing.T) {
	tests := []struct {
		name    string
		o       Map
		want    []byte
		wantErr bool
	}{
		{
			"a object",
			&encodeteststruct{
				foo: 10,
				bar: "Hello, World!",
			},
			[]byte{
				0x82, 0xA3, 0x66, 0x6F, 0x6F, 0x0A, 0xA3, 0x62,
				0x61, 0x72, 0xAD, 0x48, 0x65, 0x6C, 0x6C, 0x6F,
				0x2C, 0x20, 0x57, 0x6F, 0x72, 0x6C, 0x64, 0x21,
			},
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			buf := bytes.NewBuffer([]byte{})
			e := NewEncoder(buf)
			if err := e.EncodeMap(tt.o); (err != nil) != tt.wantErr {
				t.Errorf("EncodeMap() error = %v, wantErr %v", err, tt.wantErr)
			}
			if !bytes.Equal(buf.Bytes(), tt.want) {
				t.Errorf("EncodeMap() got = [% X], want = [% X]", buf.Bytes(), tt.want)
			}
		})
	}
}
