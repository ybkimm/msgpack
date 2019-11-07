package msgpack

import "math"

func MarshalFloat64(v float64) ([]byte, error) {
	return NewEncoder(nil).encodeFloat64(v)
}

func (e *Encoder) PutFloat64(v float64) {
	e.encodeFloat64(v)
}

func (e *Encoder) PutFloat64Key(key string, v float64) {
	e.encodeString(key)
	e.encodeFloat64(v)
}

func (e *Encoder) encodeFloat64(v float64) ([]byte, error) {
	e.grow(5)
	e.writeByte(Float64)
	e.writeUint64(math.Float64bits(v))

	return e.buf, e.err
}
