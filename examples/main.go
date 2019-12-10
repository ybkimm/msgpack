package main

import (
	"fmt"

	"github.com/ybkimm/mono/packages/encoding/msgpack"
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

// UserAccess contains some info about user access.
type UserAccess struct {
	id int
	ip string
	ua string
}

// MarshalMsgpackMap implements msgpack.Map
func (b *UserAccess) MarshalMsgpackMap(e *msgpack.Encoder, key string) error {
	switch key {
	case "id":
		return e.PutInt(b.id)
	case "ip":
		return e.PutString(b.ip)
	case "ua":
		return e.PutString(b.ua)
	}
	return nil
}

// UnmarshalMsgpackMap implements msgpack.Map
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

// Fields retuns its field names.
// it implements msgpack.Map
func (b *UserAccess) Fields() []string {
	return []string{"id", "ip", "ua"}
}
