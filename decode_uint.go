package msgpack

func UnmarshalUint(data []byte, v *uint) error {
	return NewBytesDecoder(data).DecodeUint(v)
}

func (d *Decoder) DecodeUint(v *uint) error {
	if v == nil {
		return ErrDecodeNil
	}
	return d.decodeUint(&v)
}

func UnmarshalNullableUint(data []byte, v **uint) error {
	return NewBytesDecoder(data).DecodeNullableUint(v)
}

func (d *Decoder) DecodeNullableUint(v **uint) error {
	if v == nil {
		return ErrDecodeNil
	}
	if *v != nil {
		*v = nil
	}
	return d.decodeUint(v)
}

func (d *Decoder) decodeUint(v **uint) error {
	c, err := d.nextByte()
	if err != nil {
		return err
	}

	if *v == nil {
		if c != Nil {
			*v = new(uint)
		} else {
			return nil
		}
	}

	// Positive fixnum or Negative fixnum
	if c>>7 == 0x00 || c>>5 == 0x07 {
		**v = uint(c)
		return nil
	}

	switch c {
	case Uint8:
		n, err := d.nextByte()
		if err != nil {
			return err
		}
		**v = uint(n)

	case Uint16:
		n, err := d.nextUint16()
		if err != nil {
			return err
		}
		**v = uint(n)

	case Uint32:
		n, err := d.nextUint32()
		if err != nil {
			return err
		}
		**v = uint(n)

	case Uint64:
		n, err := d.nextUint64()
		if err != nil {
			return err
		}
		**v = uint(n)
	}

	return nil
}
