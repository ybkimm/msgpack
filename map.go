package msgpack

type Map interface {
	MarshalMsgpackMap(e *Encoder)
	UnmarshalMsgpackMap(d *Decoder, key string) error
	KeySize() uint32
}

type NullableMap interface {
	Nullable
	Map
}
