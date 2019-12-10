package msgpack

// MarshalExtension returns extension value as msgpack format.
func MarshalExtension(ext Extension) ([]byte, error) {
	return NewEncoder(nil).encodeExtension(ext)
}

// PutExtension puts extension variable to encoder.
func (e *Encoder) PutExtension(ext Extension) (err error) {
	e.encodeKey()
	_, err = e.encodeExtension(ext)
	return
}

func (e *Encoder) encodeExtension(ext Extension) ([]byte, error) {
	typ := ext.ExtensionType()
	data := ext.MarshalMsgpackExtension()

	datalen := len(data)
	switch {
	case datalen == 1:
		e.grow(3)
		e.writeByte(Fixext1)

	case datalen == 2:
		e.grow(4)
		e.writeByte(Fixext2)

	case datalen == 4:
		e.grow(6)
		e.writeByte(Fixext4)

	case datalen == 8:
		e.grow(10)
		e.writeByte(Fixext8)

	case datalen == 16:
		e.grow(18)
		e.writeByte(Fixext16)

	case datalen <= ext8MaxLen:
		e.grow(datalen + 3)
		e.writeByte(Ext8)
		e.writeByte(byte(datalen))

	case datalen <= ext16MaxLen:
		e.grow(datalen + 4)
		e.writeByte(Ext16)
		e.writeUint16(uint16(datalen))

	case datalen <= ext32MaxLen:
		e.grow(datalen + 6)
		e.writeByte(Ext32)
		e.writeUint32(uint32(datalen))

	default:
		panic(ErrTooBigExtension)
	}

	e.writeByte(uint8(typ))
	e.writeBytes(data)

	return e.buf, e.err
}
