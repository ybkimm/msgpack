package msgpack

// Map represents key-value pairs of objects
type Map interface {
	MarshalMsgpackMap(e *Encoder, key string) error
	UnmarshalMsgpackMap(d *Decoder, key string) error
	Fields() []string
}

// NullableMap is a map that can be nil.
type NullableMap interface {
	Nullable
	Map
}
