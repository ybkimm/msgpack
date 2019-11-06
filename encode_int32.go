package msgpack

func (e *Encoder) EncodeInt32(v int32) error {
	e.putInt32(v)
	return e.flush()
}

func (e *Encoder) PutInt32(v int32) {
	e.putInt32(v)
}

func (e *Encoder) PutInt32Key(key string, v int32) {
	e.putString(key)
	e.putInt32(v)
}

func (e *Encoder) putInt32(v int32) {
	e.grow(5)
	e.writeByte(Int32)
	e.writeUint32(uint32(v))
}
