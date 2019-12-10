package msgpack

// MarshalBinary returns binary value as msgpack format.
func MarshalBinary(v []byte) ([]byte, error) {
	return NewEncoder(nil).encodeBinary(v)
}

// PutBinary puts binary variable to encoder.
func (e *Encoder) PutBinary(v []byte) (err error) {
	e.encodeKey()
	_, err = e.encodeBinary(v)
	return
}

func (e *Encoder) encodeBinary(v []byte) ([]byte, error) {
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
	return e.buf, e.err
}
