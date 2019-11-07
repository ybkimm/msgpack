package msgpack

import (
	"bytes"
	"testing"
)

func TestEncoder_encodeString(t *testing.T) {
	tests := []struct {
		name    string
		v       string
		want    []byte
		wantErr bool
	}{
		{
			"fixstr",
			"test",
			[]byte{0xA4, 0x74, 0x65, 0x73, 0x74},
			false,
		},
		{
			"a paragraph",
			TestString,
			TestStringData,
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := NewEncoder(nil).encodeString(tt.v)
			if (err != nil) != tt.wantErr {
				t.Errorf("encodeString() error = %v, wantErr %v", err, tt.wantErr)
			}
			if !bytes.Equal(got, tt.want) {
				t.Errorf("encodeString() got = [% X], want [% X]", got, tt.want)
			}
		})
	}
}
