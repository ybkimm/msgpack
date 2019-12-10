package msgpack

// MarshalUint32 returns uint32 value as msgpack format.
func MarshalUint32(v uint32) ([]byte, error) {
	return NewEncoder(nil).encodeUint32(v)
}

// PutUint32 puts uint32 variable to encoder.
func (e *Encoder) PutUint32(v uint32) (err error) {
	e.encodeKey()
	_, err = e.encodeUint32(v)
	return
}

func (e *Encoder) encodeUint32(v uint32) ([]byte, error) {
	e.grow(5)
	e.writeByte(Uint32)
	e.writeUint32(v)
	return e.buf, e.err
}
