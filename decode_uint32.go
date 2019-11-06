package msgpack

func (d *Decoder) DecodeUint32(v *uint32) error {
	if v == nil {
		return ErrDecodeNil
	}
	return d.decodeUint32(&v)
}

func (d *Decoder) DecodeNullableUint32(v **uint32) error {
	if v == nil {
		return ErrDecodeNil
	}
	if *v != nil {
		*v = nil
	}
	return d.decodeUint32(v)
}

func (d *Decoder) decodeUint32(v **uint32) error {
	c, err := d.nextByte()
	if err != nil {
		return err
	}

	if *v == nil {
		if c != Nil {
			*v = new(uint32)
		} else {
			return nil
		}
	}

	// Positive fixnum
	if c>>7 == 0x00 {
		**v = uint32(c)
		return nil
	}

	if c != Uint32 {
		return d.unexpectedByteErr(c)
	}

	n, err := d.nextUint32()
	if err != nil {
		return err
	}

	**v = n
	return nil
}
