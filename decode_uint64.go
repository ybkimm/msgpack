package msgpack

func (d *Decoder) DecodeUint64(v *uint64) error {
	if v == nil {
		return ErrDecodeNil
	}
	return d.decodeUint64(&v)
}

func (d *Decoder) DecodeNullableUint64(v **uint64) error {
	if v == nil {
		return ErrDecodeNil
	}
	if *v != nil {
		*v = nil
	}
	return d.decodeUint64(v)
}

func (d *Decoder) decodeUint64(v **uint64) error {
	c, err := d.nextByte()
	if err != nil {
		return err
	}

	if *v == nil {
		if c != Nil {
			*v = new(uint64)
		} else {
			return nil
		}
	}

	// Positive fixnum
	if c>>7 == 0x00 {
		**v = uint64(c)
		return nil
	}

	if c != Uint64 {
		return d.unexpectedByteErr(c)
	}

	n, err := d.nextUint64()
	if err != nil {
		return err
	}

	**v = n
	return nil
}
