package msgpack

func MarshalBool(v bool) ([]byte, error) {
	return NewEncoder(nil).encodeBool(v)
}

func (e *Encoder) PutBool(v bool) {
	e.encodeBool(v)
}

func (e *Encoder) PutBoolKey(key string, v bool) {
	e.encodeString(key)
	e.encodeBool(v)
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
