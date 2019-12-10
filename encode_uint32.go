package msgpack

func MarshalUint32(v uint32) ([]byte, error) {
	return NewEncoder(nil).encodeUint32(v)
}

func (e *Encoder) PutUint32(v uint32) {
	e.encodeUint32(v)
}

func (e *Encoder) PutUint32Key(key string, v uint32) {
	e.encodeString(key)
	e.encodeUint32(v)
}

func (e *Encoder) encodeUint32(v uint32) ([]byte, error) {
	e.grow(5)
	e.writeByte(Uint32)
	e.writeUint32(v)
	return e.buf, e.err
}
