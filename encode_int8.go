package msgpack

// MarshalInt8 returns int8 value as msgpack format.
func MarshalInt8(v int8) ([]byte, error) {
	return NewEncoder(nil).encodeInt8(v)
}

// PutInt8 puts int8 variable to encoder.
func (e *Encoder) PutInt8(v int8) (err error) {
	e.encodeKey()
	_, err = e.encodeInt8(v)
	return
}

func (e *Encoder) encodeInt8(v int8) ([]byte, error) {
	e.grow(2)
	e.writeByte(Int8)
	e.writeByte(byte(v))

	return e.buf, e.err
}
