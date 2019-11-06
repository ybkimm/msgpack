package msgpack

func (d *Decoder) DecodeInt8(v *int8) error {
	if v == nil {
		return ErrDecodeNil
	}
	return d.decodeInt8(&v)
}

func (d *Decoder) DecodeNullableInt8(v **int8) error {
	if v == nil {
		return ErrDecodeNil
	}
	if *v != nil {
		*v = nil
	}
	return d.decodeInt8(v)
}

func (d *Decoder) decodeInt8(v **int8) error {
	c, err := d.nextByte()
	if err != nil {
		return err
	}

	if *v == nil {
		if c != Nil {
			*v = new(int8)
		} else {
			return nil
		}
	}

	// Positive fixnum or Negative fixnum
	if c>>7 == 0x00 || c>>5 == 0x07 {
		**v = int8(c)
		return nil
	}

	if c != Int8 {
		return d.unexpectedByteErr(c)
	}

	n, err := d.nextByte()
	if err != nil {
		return err
	}

	**v = int8(n)
	return nil
}
