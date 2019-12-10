package msgpack

func MarshalUint16(v uint16) ([]byte, error) {
	return NewEncoder(nil).encodeUint16(v)
}

func (e *Encoder) PutUint16(v uint16) {
	e.encodeUint16(v)
}

func (e *Encoder) PutUint16Key(key string, v uint16) {
	e.encodeString(key)
	e.encodeUint16(v)
}

func (e *Encoder) encodeUint16(v uint16) ([]byte, error) {
	e.grow(3)
	e.writeByte(Uint16)
	e.writeUint16(v)
	return e.buf, e.err
}
