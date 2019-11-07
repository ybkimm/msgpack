package msgpack

import (
	"bytes"
	"testing"
)

func TestEncoder_EncodeMap(t *testing.T) {
	tests := []struct {
		name    string
		o       Map
		want    []byte
		wantErr bool
	}{
		{
			"a object",
			SmallStruct,
			TestStructData,
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := NewEncoder(nil).encodeMap(tt.o)
			if (err != nil) != tt.wantErr {
				t.Errorf("EncodeMap() error = %v, wantErr %v", err, tt.wantErr)
			}
			if !bytes.Equal(got, tt.want) {
				t.Errorf("EncodeMap() got = [% X], want = [% X]", got, tt.want)
			}
		})
	}
}
