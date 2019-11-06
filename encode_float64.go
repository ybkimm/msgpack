package msgpack

import "math"

func (e *Encoder) EncodeFloat64(v float64) error {
	e.putFloat64(v)
	return e.flush()
}

func (e *Encoder) PutFloat64(v float64) {
	e.putFloat64(v)
}

func (e *Encoder) PutFloat64Key(key string, v float64) {
	e.putString(key)
	e.putFloat64(v)
}

func (e *Encoder) putFloat64(v float64) {
	e.grow(5)
	e.writeByte(Float64)
	e.writeUint64(math.Float64bits(v))
}
