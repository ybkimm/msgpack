package msgpack

func (e *Encoder) EncodeUint32(v uint32) error {
	e.putUint32(v)
	return e.flush()
}

func (e *Encoder) PutUint32(v uint32) {
	e.putUint32(v)
}

func (e *Encoder) PutUint32Key(key string, v uint32) {
	e.putString(key)
	e.putUint32(v)
}

func (e *Encoder) putUint32(v uint32) {
	e.grow(5)
	e.writeByte(Uint32)
	e.writeUint32(v)
}
