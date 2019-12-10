package msgpack

func MarshalInt32(v int32) ([]byte, error) {
	return NewEncoder(nil).encodeInt32(v)
}

func (e *Encoder) PutInt32(v int32) {
	e.encodeInt32(v)
}

func (e *Encoder) PutInt32Key(key string, v int32) {
	e.encodeString(key)
	e.encodeInt32(v)
}

func (e *Encoder) encodeInt32(v int32) ([]byte, error) {
	e.grow(5)
	e.writeByte(Int32)
	e.writeUint32(uint32(v))

	return e.buf, e.err
}
