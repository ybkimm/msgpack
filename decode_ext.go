package msgpack

func UnmarshalExtension(data []byte, v Extension) error {
	return NewBytesDecoder(data).DecodeExtension(v)
}

func (d *Decoder) DecodeExtension(v Extension) error {
	return d.decodeExtension(v)
}

func UnmarshalNullableExtension(data []byte, v NullableExtension) error {
	return NewBytesDecoder(data).DecodeNullableExtension(v)
}

func (d *Decoder) DecodeNullableExtension(v NullableExtension) error {
	return d.DecodeExtension(v)
}

func (d *Decoder) decodeExtensionHeader(c byte) (int, int8, error) {
	var dataLen int

	switch c {
	case Fixext1:
		dataLen = 1

	case Fixext2:
		dataLen = 2

	case Fixext4:
		dataLen = 4

	case Fixext8:
		dataLen = 8

	case Fixext16:
		dataLen = 16

	case Ext8:
		n, err := d.nextByte()
		if err != nil {
			return 0, 0, err
		}
		dataLen = int(n)

	case Ext16:
		n, err := d.nextUint16()
		if err != nil {
			return 0, 0, err
		}
		dataLen = int(n)

	case Ext32:
		n, err := d.nextUint32()
		if err != nil {
			return 0, 0, err
		}
		dataLen = int(n)

	default:
		return 0, 0, d.unexpectedByteErr(c)
	}

	extType, err := d.nextByte()
	if err != nil {
		return 0, 0, err
	}

	return dataLen, int8(extType), nil
}

func (d *Decoder) decodeExtension(v Extension) error {
	if v == nil {
		return ErrDecodeNil
	}

	c, err := d.nextByte()
	if err != nil {
		return err
	}

	if c == Nil {
		if ne, ok := v.(NullableExtension); ok {
			return ne.UnmarshalMsgpackNull()
		} else {
			return d.unexpectedByteErr(c)
		}
	}

	dataLen, extType, err := d.decodeExtensionHeader(c)
	if err != nil {
		return err
	}

	if v.ExtensionType() != extType {
		return &ErrUnexpectedExtensionType{extType}
	}

	data, err := d.nextByteN(dataLen)
	if err != nil {
		return err
	}

	return v.UnmarshalMsgpackExtension(data)
}
