package msgpack

import (
	"bytes"
	"testing"
	"time"
)

func TestDecoder_DecodeTime(t *testing.T) {
	tests := []struct {
		name    string
		data    []byte
		want    time.Time
		wantErr bool
	}{
		{
			"a time",
			[]byte{0xD6, 0xFF, 0x36, 0xD9, 0xD8, 0x80},
			time.Date(1999, 3, 1, 0, 0, 0, 0, time.UTC),
			false,
		},
		{
			"a time with ns",
			[]byte{
				0xD7, 0xFF, 0x00, 0x00, 0x00, 0x04, 0x36, 0xD9,
				0xD8, 0x80,
			},
			time.Date(1999, 3, 1, 0, 0, 0, 1, time.UTC),
			false,
		},
		{
			"future",
			[]byte{
				0xC7, 0x0C, 0xFF, 0x00, 0x00, 0x00, 0x00, 0x00,
				0x00, 0x00, 0x0E, 0xE8, 0xBB, 0x98, 0x00,
			},
			time.Date(3999, 3, 1, 0, 0, 0, 0, time.UTC),
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var got time.Time
			d := NewDecoder(bytes.NewReader(tt.data))
			if err := d.DecodeTime(&got); (err != nil) != tt.wantErr {
				t.Errorf("DecodeTime() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !got.Equal(tt.want) {
				t.Errorf("DecodeTime() got = %v, want = %v", got, tt.want)
			}
		})
	}
}
