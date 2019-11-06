package msgpack

import (
	"bytes"
	"testing"
)

func TestDecoder_DecodeBinary(t *testing.T) {
	tests := []struct {
		name    string
		data    []byte
		want    []byte
		wantErr bool
	}{
		{
			"a binary",
			[]byte{0xC4, 0x04, 0x11, 0x22, 0x33, 0x44},
			[]byte{0x11, 0x22, 0x33, 0x44},
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var got []byte
			d := NewDecoder(bytes.NewReader(tt.data))
			if err := d.DecodeBinary(&got); (err != nil) != tt.wantErr {
				t.Errorf("DecodeBinary() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !bytes.Equal(got, tt.want) {
				t.Errorf("DecodeBinary() got = %v, want = %v", got, tt.want)
			}
		})
	}
}
