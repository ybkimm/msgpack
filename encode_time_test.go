package msgpack

import (
	"bytes"
	"testing"
	"time"
)

func TestEncoder_EncodeTime(t *testing.T) {
	tests := []struct {
		name    string
		t       time.Time
		want    []byte
		wantErr bool
	}{
		{
			"a local time",
			time.Date(1999, 3, 1, 9, 0, 0, 0, time.FixedZone("KST", 9*3600)),
			[]byte{0xD6, 0xFF, 0x36, 0xD9, 0xD8, 0x80},
			false,
		},
		{
			"a time with ns",
			time.Date(1999, 3, 1, 0, 0, 0, 1, time.UTC),
			[]byte{
				0xD7, 0xFF, 0x00, 0x00, 0x00, 0x04, 0x36, 0xD9,
				0xD8, 0x80,
			},
			false,
		},
		{
			"future",
			time.Date(3999, 3, 1, 0, 0, 0, 0, time.UTC),
			[]byte{
				0xC7, 0x0C, 0xFF, 0x00, 0x00, 0x00, 0x00, 0x00,
				0x00, 0x00, 0x0E, 0xE8, 0xBB, 0x98, 0x00,
			},
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := NewEncoder(nil).encodeTime(tt.t)
			if (err != nil) != tt.wantErr {
				t.Errorf("EncodeTime() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !bytes.Equal(got, tt.want) {
				t.Errorf("EncodeTime() got = [% X], want = [% X]", got, tt.want)
			}
		})
	}
}
