package msgpack

import (
	"bytes"
	"testing"
)

func TestEncoder_encodeBool(t *testing.T) {
	tests := []struct {
		name    string
		v       bool
		want    []byte
		wantErr bool
	}{
		{
			"true",
			true,
			[]byte{0xC3},
			false,
		},
		{
			"false",
			false,
			[]byte{0xC2},
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := NewEncoder(nil).encodeBool(tt.v)
			if (err != nil) != tt.wantErr {
				t.Errorf("encodeBool() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !bytes.Equal(got, tt.want) {
				t.Errorf("encodeBool() got = %v, want %v", got, tt.want)
			}
		})
	}
}
