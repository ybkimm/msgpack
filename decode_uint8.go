package msgpack

func UnmarshalUint8(data []byte, v *uint8) error {
	return NewBytesDecoder(data).DecodeUint8(v)
}

func (d *Decoder) DecodeUint8(v *uint8) error {
	if v == nil {
		return ErrDecodeNil
	}
	return d.decodeUint8(&v)
}

func UnmarshalNullableUint8(data []byte, v **uint8) error {
	return NewBytesDecoder(data).DecodeNullableUint8(v)
}

func (d *Decoder) DecodeNullableUint8(v **uint8) error {
	if v == nil {
		return ErrDecodeNil
	}
	if *v != nil {
		*v = nil
	}
	return d.decodeUint8(v)
}

func (d *Decoder) decodeUint8(v **uint8) error {
	c, err := d.nextByte()
	if err != nil {
		return err
	}

	if *v == nil {
		if c != Nil {
			*v = new(uint8)
		} else {
			return nil
		}
	}

	// Positive fixnum
	if c>>7 == 0x00 {
		**v = c
		return nil
	}

	if c != Uint8 {
		return d.unexpectedByteErr(c)
	}

	n, err := d.nextByte()
	if err != nil {
		return err
	}

	**v = n
	return nil
}
