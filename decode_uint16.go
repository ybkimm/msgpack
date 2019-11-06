package msgpack

func (d *Decoder) DecodeUint16(v *uint16) error {
	if v == nil {
		return ErrDecodeNil
	}
	return d.decodeUint16(&v)
}

func (d *Decoder) DecodeNullableUint16(v **uint16) error {
	if v == nil {
		return ErrDecodeNil
	}
	if *v != nil {
		*v = nil
	}
	return d.decodeUint16(v)
}

func (d *Decoder) decodeUint16(v **uint16) error {
	c, err := d.nextByte()
	if err != nil {
		return err
	}

	if *v == nil {
		if c != Nil {
			*v = new(uint16)
		} else {
			return nil
		}
	}

	// Positive fixnum
	if c>>7 == 0x00 {
		**v = uint16(c)
		return nil
	}

	if c != Uint16 {
		return d.unexpectedByteErr(c)
	}

	n, err := d.nextUint16()
	if err != nil {
		return err
	}

	**v = n
	return nil
}
