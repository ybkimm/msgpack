package msgpack

// Nullable represents an object that can be nil.
type Nullable interface {
	UnmarshalMsgpackNull() error
}
