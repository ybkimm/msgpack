package msgpack

func UnmarshalInt16(data []byte, v *int16) error {
	return NewBytesDecoder(data).DecodeInt16(v)
}

func (d *Decoder) DecodeInt16(v *int16) error {
	if v == nil {
		return ErrDecodeNil
	}
	return d.decodeInt16(&v)
}

func UnmarshalNullableInt16(data []byte, v **int16) error {
	return NewBytesDecoder(data).DecodeNullableInt16(v)
}

func (d *Decoder) DecodeNullableInt16(v **int16) error {
	if v == nil {
		return ErrDecodeNil
	}
	if *v != nil {
		*v = nil
	}
	return d.decodeInt16(v)
}

func (d *Decoder) decodeInt16(v **int16) error {
	c, err := d.nextByte()
	if err != nil {
		return err
	}

	if *v == nil {
		if c != Nil {
			*v = new(int16)
		} else {
			return nil
		}
	}

	// Positive fixnum or Negative fixnum
	if c>>7 == 0x00 || c>>5 == 0x07 {
		**v = int16(int8(c))
		return nil
	}

	if c != Int16 {
		return d.unexpectedByteErr(c)
	}

	n, err := d.nextUint16()
	if err != nil {
		return err
	}

	**v = int16(n)
	return nil
}
