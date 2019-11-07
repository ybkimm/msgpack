package msgpack

import (
	"bytes"
	"testing"
)

type extEncodeArgs struct {
	typ  int8
	data []byte
}

func (e extEncodeArgs) ExtensionType() int8 {
	return e.typ
}

func (e extEncodeArgs) MarshalMsgpackExtension() []byte {
	return e.data
}

func (e extEncodeArgs) UnmarshalMsgpackExtension(p []byte) error {
	panic("implement me")
}

func TestEncoder_encodeExtension(t *testing.T) {
	tests := []struct {
		name    string
		args    extEncodeArgs
		want    []byte
		wantErr bool
	}{
		{
			"a fixext1",
			extEncodeArgs{1, []byte{0x12}},
			[]byte{0xD4, 0x01, 0x12},
			false,
		},
		{
			"a fixext2",
			extEncodeArgs{1, []byte{0x12, 0x34}},
			[]byte{0xD5, 0x01, 0x12, 0x34},
			false,
		},
		{
			"a fixext4",
			extEncodeArgs{1, []byte{0x12, 0x34, 0x56, 0x78}},
			[]byte{0xD6, 0x01, 0x12, 0x34, 0x56, 0x78},
			false,
		},
		{
			"a fixext8",
			extEncodeArgs{1, []byte{
				0x12, 0x34, 0x56, 0x78, 0x9A, 0xBC, 0xDE, 0xF1,
			}},
			[]byte{
				0xD7, 0x01, 0x12, 0x34, 0x56, 0x78, 0x9A, 0xBC,
				0xDE, 0xF1,
			},
			false,
		},
		{
			"a fixext16",
			extEncodeArgs{1, []byte{
				0x12, 0x34, 0x56, 0x78, 0x9A, 0xBC, 0xDE, 0xF1,
				0x23, 0x45, 0x67, 0x89, 0xAB, 0xCD, 0xEF, 0x12,
			}},
			[]byte{
				0xD8, 0x01, 0x12, 0x34, 0x56, 0x78, 0x9A, 0xBC,
				0xDE, 0xF1, 0x23, 0x45, 0x67, 0x89, 0xAB, 0xCD,
				0xEF, 0x12,
			},
			false,
		},
		{
			"a ext8",
			extEncodeArgs{1, []byte{0x12, 0x34, 0x56}},
			[]byte{0xC7, 0x03, 0x01, 0x12, 0x34, 0x56},
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := NewEncoder(nil).encodeExtension(tt.args)
			if (err != nil) != tt.wantErr {
				t.Errorf("encodeExtension() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !bytes.Equal(got, tt.want) {
				t.Errorf("encodeExtension() got = [% X], want = [% X]", got, tt.want)
			}
		})
	}
}
