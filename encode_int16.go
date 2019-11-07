package msgpack

func MarshalInt16(v int16) ([]byte, error) {
	return NewEncoder(nil).encodeInt16(v)
}

func (e *Encoder) PutInt16(v int16) {
	e.encodeInt16(v)
}

func (e *Encoder) PutInt16Key(key string, v int16) {
	e.encodeString(key)
	e.encodeInt16(v)
}

func (e *Encoder) encodeInt16(v int16) ([]byte, error) {
	e.grow(3)
	e.writeByte(Int16)
	e.writeUint16(uint16(v))

	return e.buf, e.err
}
