package msgpack

func MarshalUint64(v uint64) ([]byte, error) {
	return NewEncoder(nil).encodeUint64(v)
}

func (e *Encoder) PutUint64(v uint64) {
	e.encodeUint64(v)
}

func (e *Encoder) PutUint64Key(key string, v uint64) {
	e.encodeString(key)
	e.encodeUint64(v)
}

func (e *Encoder) encodeUint64(v uint64) ([]byte, error) {
	e.grow(9)
	e.writeByte(Uint64)
	e.writeUint64(v)

	return e.buf, e.err
}
