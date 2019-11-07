package msgpack

func MarshalUint8(v uint8) ([]byte, error) {
	return NewEncoder(nil).encodeUint8(v)
}

func (e *Encoder) PutUint8(v uint8) {
	e.encodeUint8(v)
}

func (e *Encoder) PutUint8Key(key string, v uint8) {
	e.encodeString(key)
	e.encodeUint8(v)
}

func (e *Encoder) encodeUint8(v uint8) ([]byte, error) {
	e.grow(2)
	e.writeBytes([]byte{Uint8, byte(v)})

	return e.buf, e.err
}
