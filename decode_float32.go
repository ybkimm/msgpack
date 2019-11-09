package msgpack

import "math"

func UnmarshalFloat32(data []byte, v *float32) error {
	return NewBytesDecoder(data).DecodeFloat32(v)
}

func (d *Decoder) DecodeFloat32(v *float32) error {
	if v == nil {
		return ErrDecodeNil
	}
	return d.decodeFloat32(&v)
}

func UnmarshalNullableFloat32(data []byte, v **float32) error {
	return NewBytesDecoder(data).DecodeNullableFloat32(v)
}

func (d *Decoder) DecodeNullableFloat32(v **float32) error {
	if v == nil {
		return ErrDecodeNil
	}
	if *v != nil {
		*v = nil
	}
	return d.decodeFloat32(v)
}

func (d *Decoder) decodeFloat32(v **float32) error {
	c, err := d.nextByte()
	if err != nil {
		return err
	}

	if *v == nil {
		if c != Nil {
			*v = new(float32)
		} else {
			return nil
		}
	}

	// Positive fixnum or Negative fixnum
	if c>>7 == 0x00 || c>>5 == 0x07 {
		**v = float32(int8(c))
		return nil
	}

	if c != Float32 {
		return d.unexpectedByteErr(c)
	}

	bits, err := d.nextUint32()
	if err != nil {
		return err
	}

	**v = math.Float32frombits(bits)
	return nil
}
