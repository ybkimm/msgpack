package msgpack

func (e *Encoder) EncodeInt(v int) error {
	e.putInt(v)
	return e.flush()
}

func (e *Encoder) PutInt(v int) {
	e.putInt(v)
}

func (e *Encoder) PutIntKey(key string, v int) {
	e.putString(key)
	e.putInt(v)
}

func (e *Encoder) putInt(v int) {
	if v >= 0 {
		e.putUint(uint(v))
		return
	}

	if v >= negativeFixnumMin {
		e.grow(1)
		e.writeByte(byte(v))
		return
	}

	switch {
	case v >= -128:
		e.putInt8(int8(v))

	case v >= -32768:
		e.putInt16(int16(v))

	case v >= -2147483648:
		e.putInt32(int32(v))

	default:
		e.putInt64(int64(v))
	}
}
