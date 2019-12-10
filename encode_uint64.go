package msgpack

// MarshalUint64 returns uint64 value as msgpack format.
func MarshalUint64(v uint64) ([]byte, error) {
	return NewEncoder(nil).encodeUint64(v)
}

// PutUint64 puts uint64 variable to encoder.
func (e *Encoder) PutUint64(v uint64) (err error) {
	e.encodeKey()
	_, err = e.encodeUint64(v)
	return
}

func (e *Encoder) encodeUint64(v uint64) ([]byte, error) {
	e.grow(9)
	e.writeByte(Uint64)
	e.writeUint64(v)

	return e.buf, e.err
}
