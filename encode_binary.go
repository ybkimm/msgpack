package msgpack

func (e *Encoder) EncodeBinary(v []byte) error {
	e.putBinary(v)
	return e.flush()
}

func (e *Encoder) PutBinary(v []byte) {
	e.putBinary(v)
}

func (e *Encoder) PutBinaryKey(key string, v []byte) {
	e.putString(key)
	e.putBinary(v)
}

func (e *Encoder) putBinary(v []byte) {
	binlen := len(v)
	switch {
	case binlen <= bin8MaxLen:
		e.grow(binlen + 2)
		e.writeByte(Binary8)
		e.writeByte(uint8(binlen))

	case binlen <= bin16MaxLen:
		e.grow(binlen + 3)
		e.writeByte(Binary16)
		e.writeUint16(uint16(binlen))

	case binlen <= bin32MaxLen:
		e.grow(binlen + 5)
		e.writeByte(Binary32)
		e.writeUint32(uint32(binlen))

	default:
		panic(ErrTooLongBinary)
	}

	e.writeBytes(v)
}
