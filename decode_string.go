package msgpack

import (
	"reflect"
	"unsafe"
)

func UnmarshalString(data []byte, v *string) error {
	return NewBytesDecoder(data).DecodeString(v)
}

func (d *Decoder) DecodeString(v *string) error {
	if v == nil {
		return ErrDecodeNil
	}
	return d.decodeString(&v)
}

func UnmarshalNullableString(data []byte, v **string) error {
	return NewBytesDecoder(data).DecodeNullableString(v)
}

func (d *Decoder) DecodeNullableString(v **string) error {
	if v == nil {
		return ErrDecodeNil
	}
	if *v != nil {
		*v = nil
	}
	return d.decodeString(v)
}

func (d *Decoder) decodeStringHeader(c byte) (int, error) {
	if c>>5 == 0b0101 {
		return int(c & 0b00011111), nil
	}

	switch c {
	case String8:
		n, err := d.nextByte()
		return int(n), err

	case String16:
		n, err := d.nextUint16()
		return int(n), err

	case String32:
		n, err := d.nextUint32()
		return int(n), err

	default:
		return 0, d.unexpectedByteErr(c)
	}
}

func (d *Decoder) decodeString(v **string) error {
	c, err := d.nextByte()
	if err != nil {
		return err
	}

	if *v == nil {
		if c != Nil {
			*v = new(string)
		} else {
			return nil
		}
	}

	dataLen, err := d.decodeStringHeader(c)
	if err != nil {
		return err
	}

	data, err := d.nextByteN(dataLen)
	if err != nil {
		return err
	}

	str := reflect.StringHeader{
		Data: uintptr(unsafe.Pointer(&data[0])),
		Len:  dataLen,
	}

	**v = *(*string)(unsafe.Pointer(&str))
	return nil
}
