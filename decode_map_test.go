package msgpack

import (
	"bytes"
	"reflect"
	"testing"
)

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
			&TestStruct{},
			TestStructData,
			TestStructInstance,
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
