package msgpack

func (d *Decoder) DecodeInt32(v *int32) error {
	if v == nil {
		return ErrDecodeNil
	}
	return d.decodeInt32(&v)
}

func (d *Decoder) DecodeNullableInt32(v **int32) error {
	if v == nil {
		return ErrDecodeNil
	}
	if *v != nil {
		*v = nil
	}
	return d.decodeInt32(v)
}

func (d *Decoder) decodeInt32(v **int32) error {
	c, err := d.nextByte()
	if err != nil {
		return err
	}

	if *v == nil {
		if c != Nil {
			*v = new(int32)
		} else {
			return nil
		}
	}

	// Positive fixnum or Negative fixnum
	if c>>7 == 0x00 || c>>5 == 0x07 {
		**v = int32(int8(c))
		return nil
	}

	if c != Int32 {
		return d.unexpectedByteErr(c)
	}

	n, err := d.nextUint32()
	if err != nil {
		return err
	}

	**v = int32(n)
	return nil
}
