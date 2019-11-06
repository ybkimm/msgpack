package msgpack

import (
	"encoding/binary"
	"io"
)

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

func (d *Decoder) read() (bool, error) {
	if d.r == nil {
		return false, ErrNoReader
	}

	if len(d.buf) >= d.length {
		bufSize := d.length * 2
		if bufSize == 0 {
			bufSize = 4096
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
			return 0, io.EOF
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
	}
}
