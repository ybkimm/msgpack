package msgpack

func (e *Encoder) EncodeArray(a Array) error {
	e.putArray(a)
	return e.flush()
}

func (e *Encoder) PutArray(a Array) {
	e.putArray(a)
}

func (e *Encoder) PutArrayKey(key string, a Array) {
	e.putString(key)
	e.putArray(a)
}

func (e *Encoder) putArray(a Array) {
	e.grow(512)

	arrlen := a.Length()
	switch {
	case arrlen <= fixarrMaxLen:
		e.writeByte(fixarrPrefix | uint8(arrlen))

	case arrlen <= arr16MaxLen:
		e.writeByte(Array16)
		e.writeUint16(uint16(arrlen))

	default:
		e.writeByte(Array32)
		e.writeUint32(arrlen)
	}

	for i, l := 0, int(arrlen); i < l; i++ {
		a.MarshalMsgpackArray(e, i)
	}
}
