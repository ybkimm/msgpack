package msgpack

import (
	"bytes"
	"encoding/binary"
	"encoding/json"
	"io"
	"time"
)

func Marshal(v interface{}) ([]byte, error) {
	buf := bytes.NewBuffer(make([]byte, 0, 512))
	err := NewEncoder(buf).Encode(v)
	if err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}

type Encoder struct {
	w   io.Writer
	buf []byte
	err error
}

func NewEncoder(w io.Writer) *Encoder {
	return &Encoder{
		w: w,
	}
}

func (e *Encoder) Encode(v interface{}) error {
	if v == nil {
		e.grow(1)
		e.writeByte(Nil)
		return e.flush()
	}

	switch vv := v.(type) {
	case Map:
		return e.EncodeMap(vv)

	case Array:
		return e.EncodeArray(vv)

	case Extension:
		return e.EncodeExtension(vv)

	case bool:
		return e.EncodeBool(vv)

	case []bool:
		return e.EncodeArray((*BoolArray)(&vv))

	case int:
		return e.EncodeInt(vv)

	case []int:
		return e.EncodeArray((*IntArray)(&vv))

	case int8:
		return e.EncodeInt8(vv)

	case []int8:
		return e.EncodeArray((*Int8Array)(&vv))

	case int16:
		return e.EncodeInt16(vv)

	case []int16:
		return e.EncodeArray((*Int16Array)(&vv))

	case int32:
		return e.EncodeInt32(vv)

	case []int32:
		return e.EncodeArray((*Int32Array)(&vv))

	case int64:
		return e.EncodeInt64(vv)

	case []int64:
		return e.EncodeArray((*Int64Array)(&vv))

	case uint:
		return e.EncodeUint(vv)

	case []uint:
		return e.EncodeArray((*UintArray)(&vv))

	case uint8:
		return e.EncodeUint8(vv)

	case uint16:
		return e.EncodeUint16(vv)

	case []uint16:
		return e.EncodeArray((*Uint16Array)(&vv))

	case uint32:
		return e.EncodeUint32(vv)

	case []uint32:
		return e.EncodeArray((*Uint32Array)(&vv))

	case uint64:
		return e.EncodeUint64(vv)

	case []uint64:
		return e.EncodeArray((*Uint64Array)(&vv))

	case float32:
		return e.EncodeFloat32(vv)

	case []float32:
		return e.EncodeArray((*Float32Array)(&vv))

	case float64:
		return e.EncodeFloat64(vv)

	case []float64:
		return e.EncodeArray((*Float64Array)(&vv))

	case string:
		return e.EncodeString(vv)

	case []string:
		return e.EncodeArray((*StringArray)(&vv))

	case []byte:
		return e.EncodeBinary(vv)

	case [][]byte:
		return e.EncodeArray((*BinaryArray)(&vv))

	case time.Time:
		return e.EncodeTime(vv)

	case []time.Time:
		return e.EncodeArray((*TimeArray)(&vv))

	default:
		// Fallback: json
		src, err := json.Marshal(v)
		if err != nil {
			return err
		}
		return e.EncodeJSON(src)
	}
}

func (e *Encoder) grow(n int) {
	if cap(e.buf)-len(e.buf) > n {
		buf := make([]byte, len(e.buf), 2*cap(e.buf)+n)
		copy(buf, e.buf)
		e.buf = buf
	}
}

func (e *Encoder) writeByte(c byte) {
	e.buf = append(e.buf, c)
}

func (e *Encoder) writeBytes(p []byte) {
	e.buf = append(e.buf, p...)
}

func (e *Encoder) writeString(s string) {
	e.buf = append(e.buf, s...)
}

func (e *Encoder) writeUint16(v uint16) {
	buf := make([]byte, 2)
	binary.BigEndian.PutUint16(buf, v)
	e.writeBytes(buf)
}

func (e *Encoder) writeUint32(v uint32) {
	buf := make([]byte, 4)
	binary.BigEndian.PutUint32(buf, v)
	e.writeBytes(buf)
}

func (e *Encoder) writeUint64(v uint64) {
	buf := make([]byte, 8)
	binary.BigEndian.PutUint64(buf, v)
	e.writeBytes(buf)
}

func (e *Encoder) flush() error {
	if e.err != nil {
		return e.err
	}

	_, err := e.w.Write(e.buf)
	if err != nil {
		return err
	}

	e.buf = e.buf[:0]
	return nil
}
