package msgpack

import "math"

// MarshalFloat64 returns float64 value as msgpack format.
func MarshalFloat64(v float64) ([]byte, error) {
	return NewEncoder(nil).encodeFloat64(v)
}

// PutFloat64 puts float64 variable to encoder.
func (e *Encoder) PutFloat64(v float64) (err error) {
	e.encodeKey()
	_, err = e.encodeFloat64(v)
	return
}

func (e *Encoder) encodeFloat64(v float64) ([]byte, error) {
	e.grow(5)
	e.writeByte(Float64)
	e.writeUint64(math.Float64bits(v))

	return e.buf, e.err
}
