package msgpack

func (e *Encoder) EncodeInt16(v int16) error {
	e.putInt16(v)
	return e.flush()
}

func (e *Encoder) PutInt16(v int16) {
	e.putInt16(v)
}

func (e *Encoder) PutInt16Key(key string, v int16) {
	e.putString(key)
	e.putInt16(v)
}

func (e *Encoder) putInt16(v int16) {
	e.grow(3)
	e.writeByte(Int16)
	e.writeUint16(uint16(v))
}
