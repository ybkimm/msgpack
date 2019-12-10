package msgpack

// MarshalInt32 returns int32 value as msgpack format.
func MarshalInt32(v int32) ([]byte, error) {
	return NewEncoder(nil).encodeInt32(v)
}

// PutInt32 puts int32 variable to encoder.
func (e *Encoder) PutInt32(v int32) (err error) {
	e.encodeKey()
	_, err = e.encodeInt32(v)
	return
}

func (e *Encoder) encodeInt32(v int32) ([]byte, error) {
	e.grow(5)
	e.writeByte(Int32)
	e.writeUint32(uint32(v))

	return e.buf, e.err
}
