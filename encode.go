package msgpack

import (
	"encoding/binary"
	"encoding/json"
	"io"
	"time"
)

func Marshal(v interface{}) ([]byte, error) {
	return NewEncoder(nil).encode(v)
}

type Encoder struct {
	w   io.Writer
	buf []byte
	err error
}

func NewEncoder(w io.Writer) *Encoder {
	return (&Encoder{w: w}).Reset()
}

func (e *Encoder) encode(v interface{}) ([]byte, error) {
	if v == nil {
		e.grow(1)
		e.writeByte(Nil)
		return e.buf, e.err
	}

	switch vv := v.(type) {
	case Map:
		return e.encodeMap(vv)

	case Array:
		return e.encodeArray(vv)

	case Extension:
		return e.encodeExtension(vv)

	case bool:
		return e.encodeBool(vv)

	case []bool:
		return e.encodeArray((*BoolArray)(&vv))

	case int:
		return e.encodeInt(vv)

	case []int:
		return e.encodeArray((*IntArray)(&vv))

	case int8:
		return e.encodeInt8(vv)

	case []int8:
		return e.encodeArray((*Int8Array)(&vv))

	case int16:
		return e.encodeInt16(vv)

	case []int16:
		return e.encodeArray((*Int16Array)(&vv))

	case int32:
		return e.encodeInt32(vv)

	case []int32:
		return e.encodeArray((*Int32Array)(&vv))

	case int64:
		return e.encodeInt64(vv)

	case []int64:
		return e.encodeArray((*Int64Array)(&vv))

	case uint:
		return e.encodeUint(vv)

	case []uint:
		return e.encodeArray((*UintArray)(&vv))

	case uint8:
		return e.encodeUint8(vv)

	case uint16:
		return e.encodeUint16(vv)

	case []uint16:
		return e.encodeArray((*Uint16Array)(&vv))

	case uint32:
		return e.encodeUint32(vv)

	case []uint32:
		return e.encodeArray((*Uint32Array)(&vv))

	case uint64:
		return e.encodeUint64(vv)

	case []uint64:
		return e.encodeArray((*Uint64Array)(&vv))

	case float32:
		return e.encodeFloat32(vv)

	case []float32:
		return e.encodeArray((*Float32Array)(&vv))

	case float64:
		return e.encodeFloat64(vv)

	case []float64:
		return e.encodeArray((*Float64Array)(&vv))

	case string:
		return e.encodeString(vv)

	case []string:
		return e.encodeArray((*StringArray)(&vv))

	case []byte:
		return e.encodeBinary(vv)

	case [][]byte:
		return e.encodeArray((*BinaryArray)(&vv))

	case time.Time:
		return e.encodeTime(vv)

	case []time.Time:
		return e.encodeArray((*TimeArray)(&vv))

	default:
		// Fallback:
		src, err := json.Marshal(v)
		if err != nil {
			e.err = err
			return e.buf, e.err
		}
		return e.encodeJSON(src)
	}
}

func (e *Encoder) Reset() *Encoder {
	e.buf = make([]byte, 0, 512)
	e.err = nil
	return e
}

func (e *Encoder) grow(n int) {
	if cap(e.buf)-len(e.buf) < n {
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
