package benchmark

import (
	"github.com/francoispqt/gojay"
	"github.com/ybkimm/msgpack"
)

var _ msgpack.Map = (*benchSmallStruct)(nil)
var _ gojay.MarshalerJSONObject = (*benchSmallStruct)(nil)
var _ gojay.UnmarshalerJSONObject = (*benchSmallStruct)(nil)

var SmallStruct = &benchSmallStruct{
	id: 12345,
	ip: "192.168.0.100",
	ua: "Mozilla/5.0 (Windows NT 10.0; Win64; x64; rv:71.0) Gecko/20100101 Firefox/71.0",
}

var SmallDataJSON = []byte(`{"id":12345,"ip":"192.168.0.100","ua":"Mozilla/5.0 (Windows NT 10.0; Win64; x64; rv:71.0) Gecko/20100101 Firefox/71.0"}`)
var SmallData = []byte{
	0x83, 0xa2, 0x69, 0x64, 0xcd, 0x30, 0x39, 0xa2,
	0x69, 0x70, 0xad, 0x31, 0x39, 0x32, 0x2e, 0x31,
	0x36, 0x38, 0x2e, 0x30, 0x2e, 0x31, 0x30, 0x30,
	0xa2, 0x75, 0x61, 0xd9, 0x4e, 0x4d, 0x6f, 0x7a,
	0x69, 0x6c, 0x6c, 0x61, 0x2f, 0x35, 0x2e, 0x30,
	0x20, 0x28, 0x57, 0x69, 0x6e, 0x64, 0x6f, 0x77,
	0x73, 0x20, 0x4e, 0x54, 0x20, 0x31, 0x30, 0x2e,
	0x30, 0x3b, 0x20, 0x57, 0x69, 0x6e, 0x36, 0x34,
	0x3b, 0x20, 0x78, 0x36, 0x34, 0x3b, 0x20, 0x72,
	0x76, 0x3a, 0x37, 0x31, 0x2e, 0x30, 0x29, 0x20,
	0x47, 0x65, 0x63, 0x6b, 0x6f, 0x2f, 0x32, 0x30,
	0x31, 0x30, 0x30, 0x31, 0x30, 0x31, 0x20, 0x46,
	0x69, 0x72, 0x65, 0x66, 0x6f, 0x78, 0x2f, 0x37,
	0x31, 0x2e, 0x30,
}

type benchSmallStruct struct {
	id int
	ip string
	ua string
}

func (b *benchSmallStruct) MarshalMsgpackMap(e *msgpack.Encoder) {
	e.PutIntKey("id", b.id)
	e.PutStringKey("ip", b.ip)
	e.PutStringKey("ua", b.ua)
}

func (b *benchSmallStruct) UnmarshalMsgpackMap(d *msgpack.Decoder, key string) error {
	switch key {
	case "id":
		return d.DecodeInt(&b.id)

	case "ip":
		return d.DecodeString(&b.ip)

	case "ua":
		return d.DecodeString(&b.ua)
	}

	return nil
}

func (b *benchSmallStruct) KeySize() uint32 {
	return 3
}

func (b *benchSmallStruct) MarshalJSONObject(enc *gojay.Encoder) {
	enc.IntKey("id", b.id)
	enc.StringKey("ip", b.ip)
	enc.StringKey("ua", b.ua)
}

func (b *benchSmallStruct) IsNil() bool {
	return b == nil
}

func (b *benchSmallStruct) UnmarshalJSONObject(dec *gojay.Decoder, key string) error {
	switch key {
	case "id":
		return dec.Int(&b.id)

	case "ip":
		return dec.String(&b.ip)

	case "ua":
		return dec.String(&b.ua)
	}

	return nil
}

func (b *benchSmallStruct) NKeys() int {
	return 3
}
