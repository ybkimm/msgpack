package msgpack

func UnmarshalMap(data []byte, v Map) error {
	return NewBytesDecoder(data).DecodeMap(v)
}

func (d *Decoder) DecodeMap(v Map) error {
	return d.decodeMap(v)
}

func UnmarshalNullableMap(data []byte, v NullableMap) error {
	return NewBytesDecoder(data).DecodeNullableMap(v)
}

func (d *Decoder) DecodeNullableMap(v NullableMap) error {
	return d.decodeMap(v)
}

func (d *Decoder) decodeMapHeader(c byte) (int, error) {
	if c>>4 == 0b1000 {
		return int(c & 0b00001111), nil
	}

	switch c {
	case Map16:
		n, err := d.nextUint16()
		return int(n), err

	case Map32:
		n, err := d.nextUint32()
		return int(n), err

	default:
		return 0, d.unexpectedByteErr(c)
	}
}

func (d *Decoder) decodeMap(v Map) error {
	if v == nil {
		return ErrDecodeNil
	}

	c, err := d.nextByte()
	if err != nil {
		return err
	}

	if c == Nil {
		if no, ok := v.(NullableMap); ok {
			return no.UnmarshalMsgpackNull()
		} else {
			return d.unexpectedByteErr(c)
		}
	}

	mapLen, err := d.decodeMapHeader(c)
	if err != nil {
		return err
	}

	var key *string = new(string)

	for i := 0; i < mapLen; i++ {
		err = d.decodeString(&key)
		if err != nil {
			return err
		}

		err = v.UnmarshalMsgpackMap(d, *key)
		if err != nil {
			return err
		}
	}

	return nil
}
