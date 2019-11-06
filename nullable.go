package msgpack

type Nullable interface {
	UnmarshalMsgpackNull() error
}
