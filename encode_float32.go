package msgpack

import "math"

func (e *Encoder) EncodeFloat32(v float32) error {
	e.putFloat32(v)
	return e.flush()
}

func (e *Encoder) PutFloat32(v float32) {
	e.putFloat32(v)
}

func (e *Encoder) PutFloat32Key(key string, v float32) {
	e.putString(key)
	e.putFloat32(v)
}

func (e *Encoder) putFloat32(v float32) {
	e.grow(5)
	e.writeByte(Float32)
	e.writeUint32(math.Float32bits(v))
}
