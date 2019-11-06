package msgpack

func (e *Encoder) EncodeInt8(v int8) error {
	e.putInt8(v)
	return e.flush()
}

func (e *Encoder) PutInt8(v int8) {
	e.putInt8(v)
}

func (e *Encoder) PutInt8Key(key string, v int8) {
	e.putString(key)
	e.putInt8(v)
}

func (e *Encoder) putInt8(v int8) {
	e.grow(2)
	e.writeByte(Int8)
	e.writeByte(byte(v))
}
