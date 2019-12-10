package msgpack

import "math"

// MarshalFloat32 returns float32 value as msgpack format.
func MarshalFloat32(v float32) ([]byte, error) {
	return NewEncoder(nil).encodeFloat32(v)
}

// PutFloat32 puts float32 variable to encoder.
func (e *Encoder) PutFloat32(v float32) (err error) {
	e.encodeKey()
	_, err = e.encodeFloat32(v)
	return
}

func (e *Encoder) encodeFloat32(v float32) ([]byte, error) {
	e.grow(5)
	e.writeByte(Float32)
	e.writeUint32(math.Float32bits(v))

	return e.buf, e.err
}
