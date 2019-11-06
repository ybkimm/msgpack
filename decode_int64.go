package msgpack

func (d *Decoder) DecodeInt64(v *int64) error {
	if v == nil {
		return ErrDecodeNil
	}
	return d.decodeInt64(&v)
}

func (d *Decoder) DecodeNullableInt64(v **int64) error {
	if v == nil {
		return ErrDecodeNil
	}
	if *v != nil {
		*v = nil
	}
	return d.decodeInt64(v)
}

func (d *Decoder) decodeInt64(v **int64) error {
	c, err := d.nextByte()
	if err != nil {
		return err
	}

	if *v == nil {
		if c != Nil {
			*v = new(int64)
		} else {
			return nil
		}
	}

	// Positive fixnum or Negative fixnum
	if c>>7 == 0x00 || c>>5 == 0x07 {
		**v = int64(int8(c))
		return nil
	}

	if c != Int64 {
		return d.unexpectedByteErr(c)
	}

	n, err := d.nextUint64()
	if err != nil {
		return err
	}

	**v = int64(n)
	return nil
}
