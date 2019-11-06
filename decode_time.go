package msgpack

import "time"

func (d *Decoder) DecodeTime(v *time.Time) error {
	if v == nil {
		return ErrDecodeNil
	}
	return d.decodeTime(&v)
}

func (d *Decoder) DecodeNullableTime(v **time.Time) error {
	if v == nil {
		return ErrDecodeNil
	}
	if *v != nil {
		*v = nil
	}
	return d.decodeTime(v)
}

func (d *Decoder) decodeTime(v **time.Time) error {
	var ext Extension = &extTime{}

	if *v == nil {
		ext = &nullableExtension{ext}
	}

	if err := d.decodeExtension(ext); err != nil {
		return err
	}

	if *v == nil {
		*v = new(time.Time)
	}

	**v = unwrapNullableExtension(ext).(*extTime).Time()
	return nil
}
