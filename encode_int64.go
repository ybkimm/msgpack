package msgpack

func (e *Encoder) EncodeInt64(v int64) error {
	e.putInt64(v)
	return e.flush()
}

func (e *Encoder) PutInt64(v int64) {
	e.putInt64(v)
}

func (e *Encoder) PutInt64Key(key string, v int64) {
	e.putString(key)
	e.putInt64(v)
}

func (e *Encoder) putInt64(v int64) {
	e.grow(9)
	e.writeByte(Int64)
	e.writeUint64(uint64(v))
}
