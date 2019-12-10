package msgpack

import "math"

func MarshalFloat32(v float32) ([]byte, error) {
	return NewEncoder(nil).encodeFloat32(v)
}

func (e *Encoder) PutFloat32(v float32) {
	e.encodeFloat32(v)
}

func (e *Encoder) PutFloat32Key(key string, v float32) {
	e.encodeString(key)
	e.encodeFloat32(v)
}

func (e *Encoder) encodeFloat32(v float32) ([]byte, error) {
	e.grow(5)
	e.writeByte(Float32)
	e.writeUint32(math.Float32bits(v))

	return e.buf, e.err
}
