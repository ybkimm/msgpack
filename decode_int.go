package msgpack

func UnmarshalInt(data []byte, v *int) error {
	return NewBytesDecoder(data).DecodeInt(v)
}

func (d *Decoder) DecodeInt(v *int) error {
	if v == nil {
		return ErrDecodeNil
	}
	return d.decodeInt(&v)
}

func UnmarshalNullableInt(data []byte, v **int) error {
	return NewBytesDecoder(data).DecodeNullableInt(v)
}

func (d *Decoder) DecodeNullableInt(v **int) error {
	if v == nil {
		return ErrDecodeNil
	}
	if *v != nil {
		*v = nil
	}
	return d.decodeInt(v)
}

func (d *Decoder) decodeInt(v **int) error {
	c, err := d.nextByte()
	if err != nil {
		return err
	}

	if *v == nil {
		if c != Nil {
			*v = new(int)
		} else {
			return nil
		}
	}

	// Positive fixnum or Negative fixnum
	if c>>7 == 0x00 || c>>5 == 0x07 {
		**v = int(int8(c))
		return nil
	}

	switch c {
	case Int8:
		n, err := d.nextByte()
		if err != nil {
			return err
		}
		**v = int(int8(n))

	case Int16:
		n, err := d.nextUint16()
		if err != nil {
			return err
		}
		**v = int(int16(n))

	case Int32:
		n, err := d.nextUint32()
		if err != nil {
			return err
		}
		**v = int(int32(n))

	case Int64:
		n, err := d.nextUint64()
		if err != nil {
			return err
		}
		**v = int(int64(n))
	}

	return nil
}
