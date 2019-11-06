package msgpack

func (e *Encoder) EncodeUint64(v uint64) error {
	e.putUint64(v)
	return e.flush()
}

func (e *Encoder) PutUint64(v uint64) {
	e.putUint64(v)
}

func (e *Encoder) PutUint64Key(key string, v uint64) {
	e.putString(key)
	e.putUint64(v)
}

func (e *Encoder) putUint64(v uint64) {
	e.grow(5)
	e.writeByte(Uint64)
	e.writeUint64(v)
}
