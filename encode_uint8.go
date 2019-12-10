package msgpack

// MarshalUint8 returns uint64 value as msgpack format.
func MarshalUint8(v uint8) ([]byte, error) {
	return NewEncoder(nil).encodeUint8(v)
}

// PutUint8 puts uint8 variable to encoder.
func (e *Encoder) PutUint8(v uint8) (err error) {
	e.encodeKey()
	_, err = e.encodeUint8(v)
	return
}

func (e *Encoder) encodeUint8(v uint8) ([]byte, error) {
	e.grow(2)
	e.writeBytes([]byte{Uint8, byte(v)})

	return e.buf, e.err
}
