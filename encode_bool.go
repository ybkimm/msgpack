package msgpack

func (e *Encoder) EncodeBool(v bool) error {
	e.putBool(v)
	return e.flush()
}

func (e *Encoder) PutBool(v bool) {
	e.putBool(v)
}

func (e *Encoder) PutBoolKey(key string, v bool) {
	e.putString(key)
	e.putBool(v)
}

func (e *Encoder) putBool(v bool) {
	e.grow(1)
	if v {
		e.writeByte(True)
	} else {
		e.writeByte(False)
	}
}
