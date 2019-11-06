package msgpack

func (e *Encoder) EncodeUint16(v uint16) error {
	e.putUint16(v)
	return e.flush()
}

func (e *Encoder) PutUint16(v uint16) {
	e.putUint16(v)
}

func (e *Encoder) PutUint16Key(key string, v uint16) {
	e.putString(key)
	e.putUint16(v)
}

func (e *Encoder) putUint16(v uint16) {
	e.grow(5)
	e.writeByte(Uint16)
	e.writeUint16(v)
}
