package msgpack

// MarshalInt64 returns int64 value as msgpack format.
func MarshalInt64(v int64) ([]byte, error) {
	return NewEncoder(nil).encodeInt64(v)
}

// PutInt64 puts int64 variable to encoder.
func (e *Encoder) PutInt64(v int64) (err error) {
	e.encodeKey()
	_, err = e.encodeInt64(v)
	return
}

func (e *Encoder) encodeInt64(v int64) ([]byte, error) {
	e.grow(9)
	e.writeByte(Int64)
	e.writeUint64(uint64(v))

	return e.buf, e.err
}
