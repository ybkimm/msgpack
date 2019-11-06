package benchmark

import "github.com/ybkimm/msgpack"

var _ msgpack.Map = (*benchSmallStruct)(nil)

var SmallData = &benchSmallStruct{
	id: 12345,
	ip: "192.168.0.100",
	ua: "Mozilla/5.0 (Windows NT 10.0; Win64; x64; rv:71.0) Gecko/20100101 Firefox/71.0",
}

type benchSmallStruct struct {
	id int
	ip string
	ua string
}

func (b benchSmallStruct) MarshalMsgpackMap(e *msgpack.Encoder) {
	e.PutIntKey("id", b.id)
	e.PutStringKey("ip", b.ip)
	e.PutStringKey("ua", b.ua)
}

func (b benchSmallStruct) UnmarshalMsgpackMap(d *msgpack.Decoder, key string) error {
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

func (b benchSmallStruct) KeySize() uint32 {
	return 3
}
