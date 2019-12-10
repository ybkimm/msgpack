package msgpack

import (
	"bytes"
	"encoding/base64"
	"io"
	"math"
	"strconv"
	"time"
)

func ToJSON(data []byte) ([]byte, error) {
	d := NewDecoder(bytes.NewReader(data))

	buf := bytes.NewBuffer(make([]byte, 0, 512))
	err := d.DecodeJSON(buf)
	if err != nil {
		return nil, err
	}

	return buf.Bytes(), nil
}

func (d *Decoder) DecodeJSON(w io.Writer) error {
	return d.decodeJSON(w)
}

func (d *Decoder) decodeJSON(w io.Writer) error {
	buf := bytes.NewBuffer(make([]byte, 0, 512))

	c, err := d.nextByte()
	if err == io.EOF {
		return ErrUnexpectedEOF
	}
	if err != nil {
		return err
	}

	err = d.decodeJSONValue(buf, c)
	if err != nil {
		return err
	}

	w.Write(buf.Bytes())
	return nil
}

func (d *Decoder) decodeJSONArray(w io.Writer, c byte) error {
	arrLen, err := d.decodeArrayHeader(c)
	if err != nil {
		return err
	}

	w.Write([]byte{'['})

	for i := 0; i < arrLen; i++ {
		if i != 0 {
			w.Write([]byte{','})
		}

		c, err := d.nextByte()
		if err != nil {
			return err
		}

		err = d.decodeJSONValue(w, c)
		if err != nil {
			return err
		}
	}

	w.Write([]byte{']'})
	return nil
}

func (d *Decoder) decodeJSONMap(w io.Writer, c byte) error {
	mapLen, err := d.decodeMapHeader(c)
	if err != nil {
		return err
	}

	w.Write([]byte{'{'})

	for i := 0; i < 2*mapLen; i++ {
		c, err := d.nextByte()
		if err != nil {
			return err
		}

		if i%2 == 0 {
			if i != 0 {
				w.Write([]byte{','})
			}

			err = d.decodeJSONString(w, c)
			if err != nil {
				return err
			}

			w.Write([]byte{':'})
		} else {
			err = d.decodeJSONValue(w, c)
			if err != nil {
				return err
			}
		}

	}

	w.Write([]byte{'}'})
	return nil
}

func (d *Decoder) decodeJSONValue(w io.Writer, c byte) error {
	switch {
	case c&0b11110000 == 0b10010000, c == Array16, c == Array32:
		return d.decodeJSONArray(w, c)

	case c&0b11110000 == 0b10000000, c == Map16, c == Map32:
		return d.decodeJSONMap(w, c)

	case c == Nil:
		io.WriteString(w, "null")
		return nil

	case c == True, c == False:
		return d.decodeJSONBool(w, c)

	case c&0b10000000 == 0b00000000, c&0b11100000 == 0b11100000,
		c == Uint8, c == Uint16, c == Uint32, c == Uint64,
		c == Int8, c == Int16, c == Int32, c == Int64,
		c == Float32, c == Float64:
		return d.decodeJSONNumber(w, c)

	case c&0b11100000 == 0b10100000, c == String8, c == String16, c == String32:
		return d.decodeJSONString(w, c)

	case c == Binary8, c == Binary16, c == Binary32:
		return d.decodeJSONBinary(w, c)

	case c == Fixext1, c == Fixext2, c == Fixext4, c == Fixext8, c == Fixext16,
		c == Ext8, c == Ext16, c == Ext32:
		return d.decodeJSONExtension(w, c)

	default:
		return d.unexpectedByteErr(c)
	}
}

func (d *Decoder) decodeJSONBool(w io.Writer, c byte) error {
	switch c {
	case True:
		io.WriteString(w, "true")

	case False:
		io.WriteString(w, "false")

	default:
		return d.unexpectedByteErr(c)
	}

	return nil
}

func (d *Decoder) decodeJSONNumber(w io.Writer, c byte) error {
	if c&0b10000000 == 0b00000000 || c&0b11100000 == 0b11100000 {
		io.WriteString(w, strconv.FormatInt(int64(int8(c)), 10))
		return nil
	}

	switch c {
	case Uint8:
		n, err := d.nextByte()
		if err != nil {
			return err
		}
		io.WriteString(w, strconv.FormatUint(uint64(n), 10))

	case Uint16:
		n, err := d.nextUint16()
		if err != nil {
			return err
		}
		io.WriteString(w, strconv.FormatUint(uint64(n), 10))

	case Uint32:
		n, err := d.nextUint32()
		if err != nil {
			return err
		}
		io.WriteString(w, strconv.FormatUint(uint64(n), 10))

	case Uint64:
		n, err := d.nextUint64()
		if err != nil {
			return err
		}
		io.WriteString(w, strconv.FormatUint(n, 10))

	case Int8:
		n, err := d.nextByte()
		if err != nil {
			return err
		}
		io.WriteString(w, strconv.FormatInt(int64(int8(n)), 10))

	case Int16:
		n, err := d.nextUint16()
		if err != nil {
			return err
		}
		io.WriteString(w, strconv.FormatInt(int64(int16(n)), 10))

	case Int32:
		n, err := d.nextUint32()
		if err != nil {
			return err
		}
		io.WriteString(w, strconv.FormatInt(int64(int32(n)), 10))

	case Int64:
		n, err := d.nextUint64()
		if err != nil {
			return err
		}
		io.WriteString(w, strconv.FormatInt(int64(n), 10))

	case Float32:
		bits, err := d.nextUint32()
		if err != nil {
			return err
		}

		n := math.Float32frombits(bits)
		io.WriteString(w, strconv.FormatFloat(float64(n), 'f', -1, 32))

	case Float64:
		bits, err := d.nextUint64()
		if err != nil {
			return err
		}

		n := math.Float64frombits(bits)
		io.WriteString(w, strconv.FormatFloat(n, 'f', -1, 32))

	default:
		return d.unexpectedByteErr(c)
	}

	return nil
}

func (d *Decoder) decodeJSONString(w io.Writer, c byte) error {
	strLen, err := d.decodeStringHeader(c)
	if err != nil {
		return err
	}

	data, err := d.nextByteN(strLen)
	if err != nil {
		return err
	}

	w.Write([]byte{'"'})
	w.Write(data)
	w.Write([]byte{'"'})

	return nil
}

func (d *Decoder) decodeJSONBinary(w io.Writer, c byte) error {
	binLen, err := d.decodeBinaryHeader(c)
	if err != nil {
		return err
	}

	data, err := d.nextByteN(binLen)
	if err != nil {
		return err
	}

	w.Write([]byte{'"'})
	writeBase64Data(w, data)
	w.Write([]byte{'"'})

	return nil
}

func (d *Decoder) decodeJSONExtension(w io.Writer, c byte) error {
	dataLen, extType, err := d.decodeExtensionHeader(c)
	if err != nil {
		return err
	}

	data, err := d.nextByteN(dataLen)
	if err != nil {
		return err
	}

	switch extType {
	case -1:
		ext := &extTime{}

		err := ext.UnmarshalMsgpackExtension(data)
		if err != nil {
			return err
		}

		w.Write([]byte{'"'})
		io.WriteString(w, ext.Time().Format(time.RFC3339))
		w.Write([]byte{'"'})

	default:
		w.Write([]byte{'"'})
		writeBase64Data(w, data)
		w.Write([]byte{'"'})
	}

	return nil
}

func writeBase64Data(w io.Writer, data []byte) {
	buf := make([]byte, 0, base64.StdEncoding.EncodedLen(len(data)))
	base64.StdEncoding.Encode(buf, data)
	w.Write(buf)
}
