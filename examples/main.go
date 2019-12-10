package main

import (
	"fmt"

	"github.com/ybkimm/msgpack"
)

func main() {
	// Marshaling
	marshalResult, err := msgpack.Marshal(&UserAccess{
		id: 72,
		ip: "192.168.0.100",
		ua: "Mozilla/5.0 (Windows NT 10.0; Win64; x64; rv:71.0) Gecko/20100101 Firefox/71.0",
	})
	if err != nil {
		panic(err)
	}
	fmt.Printf("Marshal() got = [% X]\n", marshalResult)

	// Unmarshaling
	var unmarshalResult UserAccess
	err = msgpack.Unmarshal(marshalResult, &unmarshalResult)
	if err != nil {
		panic(err)
	}
	fmt.Printf("Unmarshal() got = %v\n", unmarshalResult)
}

type UserAccess struct {
	id int
	ip string
	ua string
}

func (b *UserAccess) MarshalMsgpackMap(e *msgpack.Encoder) {
	e.PutIntKey("id", b.id)
	e.PutStringKey("ip", b.ip)
	e.PutStringKey("ua", b.ua)
}

func (b *UserAccess) UnmarshalMsgpackMap(d *msgpack.Decoder, key string) error {
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

func (b *UserAccess) KeySize() uint32 {
	return 3
}
