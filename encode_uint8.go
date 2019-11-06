package msgpack

func (e *Encoder) EncodeUint8(v uint8) error {
	e.putUint8(v)
	return e.flush()
}

func (e *Encoder) PutUint8(v uint8) {
	e.putUint8(v)
}

func (e *Encoder) PutUint8Key(key string, v uint8) {
	e.putString(key)
	e.putUint8(v)
}

func (e *Encoder) putUint8(v uint8) {
	e.grow(5)
	e.writeBytes([]byte{Uint8, byte(v)})
}
