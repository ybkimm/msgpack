package msgpack

func MarshalInt64(v int64) ([]byte, error) {
	return NewEncoder(nil).encodeInt64(v)
}

func (e *Encoder) PutInt64(v int64) {
	e.encodeInt64(v)
}

func (e *Encoder) PutInt64Key(key string, v int64) {
	e.encodeString(key)
	e.encodeInt64(v)
}

func (e *Encoder) encodeInt64(v int64) ([]byte, error) {
	e.grow(9)
	e.writeByte(Int64)
	e.writeUint64(uint64(v))

	return e.buf, e.err
}
