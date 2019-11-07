package msgpack

func MarshalUint(v uint) ([]byte, error) {
	return NewEncoder(nil).encodeUint(v)
}

func (e *Encoder) PutUint(v uint) {
	e.encodeUint(v)
}

func (e *Encoder) PutUintKey(key string, v uint) {
	e.encodeString(key)
	e.encodeUint(v)
}

func (e *Encoder) encodeUint(v uint) ([]byte, error) {
	if v <= positiveFixnumMax {
		e.grow(1)
		e.writeByte(byte(v))
	} else if v>>8 == 0 {
		e.encodeUint8(uint8(v))
	} else if v>>16 == 0 {
		e.encodeUint16(uint16(v))
	} else if v>>32 == 0 {
		e.encodeUint32(uint32(v))
	} else {
		e.encodeUint64(uint64(v))
	}

	return e.buf, e.err
}
