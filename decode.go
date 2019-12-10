package msgpack

import (
	"encoding/binary"
	"io"
	"reflect"
	"runtime/debug"
	"time"
)

func Unmarshal(data []byte, out interface{}) error {
	return NewBytesDecoder(data).Decode(out)
}

type Decoder struct {
	r      io.Reader
	buf    []byte
	cursor int
	length int
}

// NewDecoder returns new decoder instance.
func NewDecoder(r io.Reader) *Decoder {
	return &Decoder{
		r: r,
	}
}

func NewBytesDecoder(data []byte) *Decoder {
	return &Decoder{
		buf:    data,
		length: len(data),
	}
}

func (d *Decoder) Decode(v interface{}) error {
	switch vv := v.(type) {
	case Map:
		return d.DecodeMap(vv)

	case Array:
		return d.DecodeArray(vv)

	case Extension:
		return d.DecodeExtension(vv)

	case *bool:
		return d.DecodeBool(vv)

	case *int:
		return d.DecodeInt(vv)

	case *int8:
		return d.DecodeInt8(vv)

	case *int16:
		return d.DecodeInt16(vv)

	case *int32:
		return d.DecodeInt32(vv)

	case *int64:
		return d.DecodeInt64(vv)

	case *uint:
		return d.DecodeUint(vv)

	case *uint8:
		return d.DecodeUint8(vv)

	case *uint16:
		return d.DecodeUint16(vv)

	case *uint32:
		return d.DecodeUint32(vv)

	case *uint64:
		return d.DecodeUint64(vv)

	case *float32:
		return d.DecodeFloat32(vv)

	case *float64:
		return d.DecodeFloat64(vv)

	case *string:
		return d.DecodeString(vv)

	case *[]byte:
		return d.DecodeBinary(vv)

	case *time.Time:
		return d.DecodeTime(vv)

	default:
		return &ErrUnsupportedType{reflect.TypeOf(v)}
	}
}

func (d *Decoder) read() (bool, error) {
	if d.r == nil {
		return false, nil
	}

	if len(d.buf) >= d.length {
		var bufSize int

		if d.length == 0 {
			bufSize = 4096
		} else {
			bufSize = d.length * 2
		}

		buf := make([]byte, bufSize, bufSize)
		copy(buf, d.buf)

		d.buf = buf
	}

	n, err := d.r.Read(d.buf[d.length:])
	if err != nil {
		if err != io.EOF {
			return false, err
		}
		if n == 0 {
			return false, nil
		}
	}

	d.length = d.length + n
	return true, nil
}

func (d *Decoder) nextByte() (byte, error) {
	if d.cursor >= d.length {
		ok, err := d.read()
		if err != nil {
			return 0, err
		}
		if !ok {
			return 0, io.ErrUnexpectedEOF
		}
	}

	c := d.buf[d.cursor]
	d.cursor++

	return c, nil
}

func (d *Decoder) nextByteN(n int) ([]byte, error) {
	if n <= 0 {
		return nil, nil
	}

	if n == 1 {
		c, err := d.nextByte()
		if err != nil {
			return nil, err
		}
		return []byte{c}, nil
	}

	for d.cursor+n > d.length {
		ok, err := d.read()
		if err != nil {
			return nil, err
		}
		if !ok {
			return nil, io.ErrUnexpectedEOF
		}
	}

	data := d.buf[d.cursor : d.cursor+n]
	d.cursor += n

	return data, nil
}

func (d *Decoder) nextUint16() (uint16, error) {
	data, err := d.nextByteN(2)
	if err != nil {
		return 0, err
	}
	return binary.BigEndian.Uint16(data), nil
}

func (d *Decoder) nextUint32() (uint32, error) {
	data, err := d.nextByteN(4)
	if err != nil {
		return 0, err
	}
	return binary.BigEndian.Uint32(data), nil
}

func (d *Decoder) nextUint64() (uint64, error) {
	data, err := d.nextByteN(8)
	if err != nil {
		return 0, err
	}
	return binary.BigEndian.Uint64(data), nil
}

func (d *Decoder) unexpectedByteErr(c byte) *ErrUnexpectedByte {
	return &ErrUnexpectedByte{
		Byte:     c,
		Position: d.cursor - 1,
		Stack:    debug.Stack(),
	}
}
