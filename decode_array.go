package msgpack

func (d *Decoder) DecodeArray(v Array) error {
	if v == nil {
		return ErrDecodeNil
	}
	return d.decodeArray(v)
}

func (d *Decoder) DecodeNullableArray(v NullableArray) error {
	return d.DecodeArray(v)
}

func (d *Decoder) decodeArrayHeader(c byte) (int, error) {
	if c&0b11110000 == 0b10010000 {
		return int(c & 0b00001111), nil
	} else {
		switch c {
		case Array16:
			n, err := d.nextUint16()
			return int(n), err

		case Array32:
			n, err := d.nextUint32()
			return int(n), err

		default:
			return 0, d.unexpectedByteErr(c)
		}
	}
}

func (d *Decoder) decodeArray(v Array) error {
	if v == nil {
		return ErrDecodeNil
	}

	c, err := d.nextByte()
	if err != nil {
		return err
	}

	if c == Nil {
		if na, ok := v.(NullableArray); ok {
			return na.UnmarshalMsgpackNull()
		} else {
			return d.unexpectedByteErr(c)
		}
	}

	arrLen, err := d.decodeArrayHeader(c)
	if err != nil {
		return err
	}

	return v.UnmarshalMsgpackArray(d, arrLen)
}
