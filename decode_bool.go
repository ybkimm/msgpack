package msgpack

func UnmarshalBool(data []byte, v *bool) error {
	return NewBytesDecoder(data).DecodeBool(v)
}

func (d *Decoder) DecodeBool(v *bool) error {
	if v == nil {
		return ErrDecodeNil
	}
	return d.decodeBool(&v)
}

func UnmarshalNullableBool(data []byte, v **bool) error {
	return NewBytesDecoder(data).DecodeNullableBool(v)
}

func (d *Decoder) DecodeNullableBool(v **bool) error {
	if v == nil {
		return ErrDecodeNil
	}
	if *v != nil {
		*v = nil
	}
	return d.decodeBool(v)
}

func (d *Decoder) decodeBool(v **bool) error {
	c, err := d.nextByte()
	if err != nil {
		return err
	}

	if *v == nil {
		if c != Nil {
			*v = new(bool)
		} else {
			return nil
		}
	}

	if c == True {
		**v = true
	} else if c == False {
		**v = false
	} else {
		return d.unexpectedByteErr(c)
	}

	return nil
}
