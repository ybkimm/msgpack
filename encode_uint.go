package msgpack

func (e *Encoder) EncodeUint(v uint) error {
	e.putUint(v)
	return e.flush()
}

func (e *Encoder) PutUint(v uint) {
	e.putUint(v)
}

func (e *Encoder) PutUintKey(key string, v uint) {
	e.putString(key)
	e.putUint(v)
}

func (e *Encoder) putUint(v uint) {
	if v <= positiveFixnumMax {
		e.grow(1)
		e.writeByte(byte(v))
		return
	}

	switch {
	case v>>8 == 0:
		e.putUint8(uint8(v))

	case v>>16 == 0:
		e.putUint16(uint16(v))

	case v>>32 == 0:
		e.putUint32(uint32(v))

	default:
		e.putUint64(uint64(v))
	}
}
