package msgpack

import (
	"bytes"
	"encoding/json"
	"io"
	"math"
	"reflect"
)

func FromJSON(data []byte) ([]byte, error) {
	buf := bytes.NewBuffer(make([]byte, 0, 512))
	e := NewEncoder(buf)
	err := e.EncodeJSON(data)
	if err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}

func (e *Encoder) EncodeJSON(data []byte) error {
	e.putJSON(data)
	return e.flush()
}

func (e *Encoder) PutJSON(data []byte) {
	e.putJSON(data)
}

func (e *Encoder) PutJSONKey(key string, data []byte) {
	e.putString(key)
	e.putJSON(data)
}

func (e *Encoder) putJSON(data []byte) {
	decoder := json.NewDecoder(bytes.NewReader(data))

	// Check first token
	t, err := decoder.Token()
	if err == io.EOF {
		return
	} else if err != nil {
		e.err = err
		return
	}

	// Convert
	if v, ok := t.(json.Delim); ok {
		if v == '{' {
			e.putJSONObject(decoder)
		} else if v == '[' {
			e.putJSONArray(decoder)
		} else {
			err = &ErrUnexpectedJSONToken{t}
		}
	} else {
		e.putJSONValue(t)
	}

	// More tokens?
	if decoder.More() {
		t, err := decoder.Token()
		if err != nil {
			e.err = err
		} else {
			e.err = &ErrUnexpectedJSONToken{Token: t}
		}
	}
}

func (e *Encoder) putJSONObject(decoder *json.Decoder) {
	var (
		buf = bytes.NewBuffer(make([]byte, 0, 512))
		je  = NewEncoder(buf)
		i   int
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
			return
		}

		// Is key?
		if isKey {
			switch v := t.(type) {
			case json.Delim:
				if v == '}' {
					break LOOP
				} else {
					e.err = &ErrUnexpectedJSONToken{t}
					return
				}

			case string:
				je.putString(v)
				i++

			default:
				e.err = &ErrUnexpectedJSONToken{t}
				return
			}
		} else {
			switch v := t.(type) {
			case json.Delim:
				if v == '{' {
					je.putJSONObject(decoder)
				} else if v == '[' {
					je.putJSONArray(decoder)
				} else {
					err = &ErrUnexpectedJSONToken{t}
					return
				}

			default:
				je.putJSONValue(v)
			}
		}
	}

	e.err = je.flush()
	buflen := buf.Len()
	switch {
	case i <= fixmapMaxLen:
		e.grow(buflen + 1)
		e.writeByte(fixmapPrefix | byte(i))

	case i <= map16MaxLen:
		e.grow(buflen + 3)
		e.writeByte(Map16)
		e.writeUint16(uint16(i))

	case i <= map32MaxLen:
		e.grow(buflen + 5)
		e.writeByte(Map32)
		e.writeUint32(uint32(i))

	default:
		e.err = ErrTooBigMap
	}

	e.writeBytes(buf.Bytes())
}

func (e *Encoder) putJSONArray(decoder *json.Decoder) {
	var (
		buf = bytes.NewBuffer(make([]byte, 0, 512))
		je  = NewEncoder(buf)
		i   int
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
			return
		}

		switch v := t.(type) {
		case json.Delim:
			if v == '{' {
				je.putJSONObject(decoder)
			} else if v == '[' {
				je.putJSONArray(decoder)
			} else if v == ']' {
				break LOOP
			} else {
				err = &ErrUnexpectedJSONToken{t}
				return
			}

		default:
			je.putJSONValue(v)
		}
	}

	e.err = je.flush()
	buflen := buf.Len()
	switch {
	case i <= fixarrMaxLen:
		e.grow(buflen + 1)
		e.writeByte(fixarrPrefix | byte(i))

	case i <= arr16MaxLen:
		e.grow(buflen + 3)
		e.writeByte(Array16)
		e.writeUint16(uint16(i))

	case i <= arr32MaxLen:
		e.grow(buflen + 5)
		e.writeByte(Array32)
		e.writeUint32(uint32(i))

	default:
		e.err = ErrTooBigMap
	}

	e.writeBytes(buf.Bytes())
}

func (e *Encoder) putJSONValue(v interface{}) {
	if v == nil {
		e.writeByte(Nil)
		return
	}

	switch vv := v.(type) {
	case bool:
		e.putBool(vv)

	case float64:
		if math.Abs(math.Mod(vv, 1)) > 0 {
			e.putFloat64(vv)
		} else if vv < 0 {
			e.putInt(int(vv))
		} else {
			e.putUint(uint(vv))
		}

	case string:
		e.putString(vv)

	default:
		e.err = &ErrUnsupportedType{reflect.TypeOf(v)}
	}
}
