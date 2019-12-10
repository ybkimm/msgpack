package msgpack

func MarshalArray(a Array) ([]byte, error) {
	return NewEncoder(nil).encodeArray(a)
}

func (e *Encoder) PutArray(a Array) {
	e.encodeArray(a)
}

func (e *Encoder) PutArrayKey(key string, a Array) {
	e.encodeString(key)
	e.encodeArray(a)
}

func (e *Encoder) encodeArray(a Array) ([]byte, error) {
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

	return e.buf, e.err
}
