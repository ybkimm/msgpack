package msgpack

import "time"

func MarshalTime(t time.Time) ([]byte, error) {
	return NewEncoder(nil).encodeTime(t)
}

func (e *Encoder) PutTime(t time.Time) {
	e.encodeTime(t.UTC())
}

func (e *Encoder) PutTimeKey(key string, t time.Time) {
	e.encodeString(key)
	e.encodeTime(t.UTC())
}

func (e *Encoder) encodeTime(t time.Time) ([]byte, error) {
	if t.Location() != time.UTC {
		return e.encodeTime(t.UTC())
	}

	var (
		seconds     = uint64(t.Unix())
		nanoseconds = uint64(t.UnixNano()) - 1000000000*seconds
	)

	return e.encodeExtension(&extTime{
		seconds:     seconds,
		nanoseconds: nanoseconds,
	})
}
