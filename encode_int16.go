package msgpack

// MarshalInt16 returns int16 value as msgpack format.
func MarshalInt16(v int16) ([]byte, error) {
	return NewEncoder(nil).encodeInt16(v)
}

// PutInt16 puts int16 variable to encoder.
func (e *Encoder) PutInt16(v int16) (err error) {
	e.encodeKey()
	_, err = e.encodeInt16(v)
	return
}

func (e *Encoder) encodeInt16(v int16) ([]byte, error) {
	e.grow(3)
	e.writeByte(Int16)
	e.writeUint16(uint16(v))

	return e.buf, e.err
}
