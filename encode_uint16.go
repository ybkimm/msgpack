package msgpack

// MarshalUint16 returns uint16 value as msgpack format.
func MarshalUint16(v uint16) ([]byte, error) {
	return NewEncoder(nil).encodeUint16(v)
}

// PutUint16 puts uint16 variable to encoder.
func (e *Encoder) PutUint16(v uint16) (err error) {
	e.encodeKey()
	_, err = e.encodeUint16(v)
	return
}

func (e *Encoder) encodeUint16(v uint16) ([]byte, error) {
	e.grow(3)
	e.writeByte(Uint16)
	e.writeUint16(v)
	return e.buf, e.err
}
