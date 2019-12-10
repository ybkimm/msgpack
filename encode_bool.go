package msgpack

// MarshalBool returns bool value as msgpack format.
func MarshalBool(v bool) ([]byte, error) {
	return NewEncoder(nil).encodeBool(v)
}

// PutBool puts bool variable to encoder.
func (e *Encoder) PutBool(v bool) (err error) {
	e.encodeKey()
	_, err = e.encodeBool(v)
	return
}

func (e *Encoder) encodeBool(v bool) ([]byte, error) {
	e.grow(1)
	if v {
		e.writeByte(True)
	} else {
		e.writeByte(False)
	}
	return e.buf, e.err
}
