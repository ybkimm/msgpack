package msgpack

import (
	"bytes"
	"testing"
)

func TestToJSON(t *testing.T) {
	tests := []struct {
		name    string
		data    []byte
		want    []byte
		wantErr bool
	}{
		{
			"simple object",
			[]byte{
				0x84, 0xA6, 0x73, 0x74, 0x72, 0x69, 0x6E, 0x67,
				0xA6, 0x73, 0x74, 0x72, 0x69, 0x6E, 0x67, 0xA6,
				0x6E, 0x75, 0x6D, 0x62, 0x65, 0x72, 0x01, 0xA4,
				0x62, 0x6F, 0x6F, 0x6C, 0xC3, 0xA3, 0x6E, 0x69,
				0x6C, 0xC0,
			},
			[]byte(`{"string":"string","number":1,"bool":true,"nil":null}`),
			false,
		},
		{
			"simple array",
			[]byte{
				0x94, 0xA6, 0x73, 0x74, 0x72, 0x69, 0x6E, 0x67,
				0x01, 0xC3, 0xC0,
			},
			[]byte(`["string",1,true,null]`),
			false,
		},
		{
			"complex object",
			[]byte{
				0x86, 0xA6, 0x73, 0x74, 0x72, 0x69, 0x6E, 0x67,
				0xA6, 0x73, 0x74, 0x72, 0x69, 0x6E, 0x67, 0xA6,
				0x6E, 0x75, 0x6D, 0x62, 0x65, 0x72, 0x01, 0xA4,
				0x62, 0x6F, 0x6F, 0x6C, 0xC3, 0xA3, 0x6E, 0x69,
				0x6C, 0xC0, 0xA9, 0x69, 0x6E, 0x6E, 0x65, 0x72,
				0x5F, 0x6F, 0x62, 0x6A, 0x82, 0xA1, 0x61, 0xA1,
				0x61, 0xA1, 0x62, 0xA1, 0x62, 0xA9, 0x69, 0x6E,
				0x6E, 0x65, 0x72, 0x5F, 0x61, 0x72, 0x72, 0x92,
				0x82, 0xA1, 0x61, 0xA1, 0x61, 0xA1, 0x62, 0xA1,
				0x62, 0x82, 0xA1, 0x61, 0xA1, 0x61, 0xA1, 0x62,
				0xA1, 0x62,
			},
			[]byte(`{"string":"string","number":1,"bool":true,"nil":null,"inner_obj":{"a":"a","b":"b"},"inner_arr":[{"a":"a","b":"b"},{"a":"a","b":"b"}]}`),
			false,
		},
		{
			"single value",
			[]byte{0x00},
			[]byte(`0`),
			false,
		},
		{
			"invalid object",
			[]byte{
				0x84, 0xA6, 0x73, 0x74, 0x72, 0x69, 0x6E, 0x67,
			},
			nil,
			true,
		},
		{
			"invalid array",
			[]byte{0x94, 0x00, 0x00, 0x00},
			nil,
			true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ToJSON(tt.data)
			if (err != nil) != tt.wantErr {
				t.Errorf("ToJSON() error = %v, wantErr %v, got = %v", err, tt.wantErr, string(got))
				return
			}
			if !bytes.Equal(got, tt.want) {
				t.Errorf("ToJSON() got = %v, want %v", string(got), string(tt.want))
			}
		})
	}
}
