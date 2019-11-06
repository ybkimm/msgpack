package msgpack

func (d *Decoder) DecodeBinary(v *[]byte) error {
	if v == nil {
		return ErrDecodeNil
	}
	return d.decodeBinary(&v)
}

func (d *Decoder) DecodeNullableBinary(v **[]byte) error {
	if v == nil {
		return ErrDecodeNil
	}
	if *v != nil {
		*v = nil
	}
	return d.decodeBinary(v)
}

func (d *Decoder) decodeBinaryHeader(c byte) (int, error) {
	switch c {
	case Binary8:
		n, err := d.nextByte()
		return int(n), err

	case Binary16:
		n, err := d.nextUint16()
		return int(n), err

	case Binary32:
		n, err := d.nextUint32()
		return int(n), err

	default:
		return 0, d.unexpectedByteErr(c)
	}
}

func (d *Decoder) decodeBinary(v **[]byte) error {
	c, err := d.nextByte()
	if err != nil {
		return err
	}

	if *v == nil {
		if c != Nil {
			*v = new([]byte)
		} else {
			return nil
		}
	}

	dataLen, err := d.decodeBinaryHeader(c)
	if err != nil {
		return err
	}

	data, err := d.nextByteN(dataLen)
	if err != nil {
		return err
	}

	**v = data
	return nil
}
