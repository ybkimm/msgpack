package msgpack

func MarshalInt8(v int8) ([]byte, error) {
	return NewEncoder(nil).encodeInt8(v)
}

func (e *Encoder) PutInt8(v int8) {
	e.encodeInt8(v)
}

func (e *Encoder) PutInt8Key(key string, v int8) {
	e.encodeString(key)
	e.encodeInt8(v)
}

func (e *Encoder) encodeInt8(v int8) ([]byte, error) {
	e.grow(2)
	e.writeByte(Int8)
	e.writeByte(byte(v))

	return e.buf, e.err
}
