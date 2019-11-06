package msgpack

import (
	"math"
)

func (d *Decoder) DecodeFloat64(v *float64) error {
	if v == nil {
		return ErrDecodeNil
	}
	return d.decodeFloat64(&v)
}

func (d *Decoder) DecodeNullableFloat64(v **float64) error {
	if v == nil {
		return ErrDecodeNil
	}
	if *v != nil {
		*v = nil
	}
	return d.decodeFloat64(v)
}

func (d *Decoder) decodeFloat64(v **float64) error {
	c, err := d.nextByte()
	if err != nil {
		return err
	}

	if *v == nil {
		if c != Nil {
			*v = new(float64)
		} else {
			return nil
		}
	}

	// Positive fixnum or Negative fixnum
	if c>>7 == 0x00 || c>>5 == 0x07 {
		**v = float64(int8(c))
		return nil
	}

	if c != Float64 {
		return d.unexpectedByteErr(c)
	}

	bits, err := d.nextUint64()
	if err != nil {
		return err
	}

	**v = math.Float64frombits(bits)
	return nil
}
