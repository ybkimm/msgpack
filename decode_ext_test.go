package msgpack

import (
	"bytes"
	"testing"
)

var _ Extension = (*decodetestext)(nil)

type decodetestext struct {
	data []byte
}

func (d *decodetestext) ExtensionType() int8 {
	return 1
}

func (d *decodetestext) MarshalMsgpackExtension() []byte {
	panic("implement me")
}

func (d *decodetestext) UnmarshalMsgpackExtension(p []byte) error {
	d.data = p
	return nil
}

func TestDecoder_DecodeExtension(t *testing.T) {
	tests := []struct {
		name    string
		data    []byte
		want    []byte
		wantErr bool
	}{
		{
			"fixext1",
			[]byte{0xD4, 0x01, 0x01},
			[]byte{0x01},
			false,
		},
		{
			"fixext2",
			[]byte{0xD5, 0x01, 0x01, 0x23},
			[]byte{0x01, 0x23},
			false,
		},
		{
			"fixext4",
			[]byte{0xD6, 0x01, 0x01, 0x23, 0x45, 0x67},
			[]byte{0x01, 0x23, 0x45, 0x67},
			false,
		},
		{
			"fixext8",
			[]byte{
				0xD7, 0x01, 0x01, 0x23, 0x45, 0x67, 0x89, 0xAB,
				0xCD, 0xEF,
			},
			[]byte{
				0x01, 0x23, 0x45, 0x67, 0x89, 0xAB, 0xCD, 0xEF,
			},
			false,
		},
		{
			"fixext16",
			[]byte{
				0xD8, 0x01, 0x01, 0x23, 0x45, 0x67, 0x89, 0xAB,
				0xCD, 0xEF, 0x01, 0x23, 0x45, 0x67, 0x89, 0xAB,
				0xCD, 0xEF,
			},
			[]byte{
				0x01, 0x23, 0x45, 0x67, 0x89, 0xAB, 0xCD, 0xEF,
				0x01, 0x23, 0x45, 0x67, 0x89, 0xAB, 0xCD, 0xEF,
			},
			false,
		},
		{
			"ext8",
			[]byte{
				0xC7, 0x0F, 0x01, 0x01, 0x23, 0x45, 0x67, 0x89,
				0xAB, 0xCD, 0xEF, 0x01, 0x23, 0x45, 0x67, 0x89,
				0xAB, 0xCD,
			},
			[]byte{
				0x01, 0x23, 0x45, 0x67, 0x89, 0xAB, 0xCD, 0xEF,
				0x01, 0x23, 0x45, 0x67, 0x89, 0xAB, 0xCD,
			},
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var got decodetestext
			d := NewDecoder(bytes.NewReader(tt.data))
			if err := d.DecodeExtension(&got); (err != nil) != tt.wantErr {
				t.Errorf("DecodeExtension() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !bytes.Equal(got.data, tt.want) {
				t.Errorf("DecodeExtension() got = %v, want = %v", got, tt.want)
			}
		})
	}
}
