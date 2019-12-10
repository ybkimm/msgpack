package msgpack

func MarshalInt(v int) ([]byte, error) {
	return NewEncoder(nil).encodeInt(v)
}

func (e *Encoder) PutInt(v int) {
	e.encodeInt(v)
}

func (e *Encoder) PutIntKey(key string, v int) {
	e.encodeString(key)
	e.encodeInt(v)
}

func (e *Encoder) encodeInt(v int) ([]byte, error) {
	if v >= 0 {
		e.encodeUint(uint(v))
		return e.buf, e.err
	}

	if v >= negativeFixnumMin {
		e.grow(1)
		e.writeByte(byte(v))
		return e.buf, e.err
	}

	switch {
	case v >= -128:
		return e.encodeInt8(int8(v))

	case v >= -32768:
		return e.encodeInt16(int16(v))

	case v >= -2147483648:
		return e.encodeInt32(int32(v))

	default:
		return e.encodeInt64(int64(v))
	}
}
