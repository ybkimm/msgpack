package msgpack

import (
	"bytes"
	"encoding/json"
	"io"
	"math"
)

func FromJSON(data []byte) ([]byte, error) {
	return NewEncoder(nil).encodeJSON(data)
}

func (e *Encoder) PutJSON(data []byte) {
	e.encodeJSON(data)
}

func (e *Encoder) PutJSONKey(key string, data []byte) {
	e.encodeString(key)
	e.encodeJSON(data)
}

func (e *Encoder) encodeJSON(data []byte) ([]byte, error) {
	decoder := json.NewDecoder(bytes.NewReader(data))

	// Check first token
	t, err := decoder.Token()
	if err == io.EOF {
		return e.buf, e.err
	}
	if err != nil {
		e.err = err
		return e.buf, e.err
	}

	return e.encodeJSONValue(decoder, t)
}

func (e *Encoder) encodeJSONObject(decoder *json.Decoder) ([]byte, error) {
	var (
		je = NewEncoder(nil)
		i  int
	)

LOOP:
	for isKey := true; ; isKey = !isKey {
		t, err := decoder.Token()
		if err != nil {
			if err == io.EOF {
				e.err = ErrUnexpectedEOF
			} else {
				e.err = err
			}
			return e.buf, e.err
		}

		// Is key?
		if isKey {
			if v, ok := t.(json.Delim); ok && v == '}' {
				break LOOP
			} else if str, ok := t.(string); ok {
				je.encodeString(str)
			} else {
				je.err = &ErrUnexpectedJSONToken{t}
			}
			i++
		} else {
			je.encodeJSONValue(decoder, t)
		}

		if je.err != nil {
			e.err = je.err
			return e.buf, e.err
		}
	}

	bufLen := len(je.buf)

	switch {
	case i <= fixmapMaxLen:
		e.grow(bufLen + 1)
		e.writeByte(fixmapPrefix | byte(i))

	case i <= map16MaxLen:
		e.grow(bufLen + 3)
		e.writeByte(Map16)
		e.writeUint16(uint16(i))

	case i <= map32MaxLen:
		e.grow(bufLen + 5)
		e.writeByte(Map32)
		e.writeUint32(uint32(i))

	default:
		e.err = ErrTooBigMap
	}

	e.writeBytes(je.buf)
	return e.buf, e.err
}

func (e *Encoder) encodeJSONArray(decoder *json.Decoder) ([]byte, error) {
	var (
		je = NewEncoder(nil)
		i  int
	)

LOOP:
	for ; ; i++ {
		t, err := decoder.Token()
		if err != nil {
			if err == io.EOF {
				e.err = ErrUnexpectedEOF
			} else {
				e.err = err
			}
			return e.buf, e.err
		}

		if v, ok := t.(json.Delim); ok && v == ']' {
			break LOOP
		} else {
			je.encodeJSONValue(decoder, t)
		}

		if je.err != nil {
			e.err = je.err
			return e.buf, e.err
		}
	}

	bufLen := len(je.buf)

	switch {
	case i <= fixarrMaxLen:
		e.grow(bufLen + 1)
		e.writeByte(fixarrPrefix | byte(i))

	case i <= arr16MaxLen:
		e.grow(bufLen + 3)
		e.writeByte(Array16)
		e.writeUint16(uint16(i))

	case i <= arr32MaxLen:
		e.grow(bufLen + 5)
		e.writeByte(Array32)
		e.writeUint32(uint32(i))

	default:
		e.err = ErrTooBigMap
	}

	e.writeBytes(je.buf)
	return e.buf, e.err
}

func (e *Encoder) encodeJSONValue(decoder *json.Decoder, v json.Token) ([]byte, error) {
	if v == nil {
		e.writeByte(Nil)
	} else {
		switch vv := v.(type) {
		case json.Delim:
			if vv == '{' {
				return e.encodeJSONObject(decoder)
			} else if vv == '[' {
				return e.encodeJSONArray(decoder)
			} else {
				e.err = &ErrUnexpectedJSONToken{vv}
			}

		case bool:
			return e.encodeBool(vv)

		case float64:
			if math.Abs(math.Mod(vv, 1)) > 0 {
				return e.encodeFloat64(vv)
			} else if vv < 0 {
				return e.encodeInt(int(vv))
			} else {
				return e.encodeUint(uint(vv))
			}

		case string:
			return e.encodeString(vv)

		default:
			e.err = &ErrUnexpectedJSONToken{v}
		}
	}

	return e.buf, e.err
}
