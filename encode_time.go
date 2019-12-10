package msgpack

import "time"

// MarshalTime returns time.Time value as msgpack format.
func MarshalTime(t time.Time) ([]byte, error) {
	return NewEncoder(nil).encodeTime(t)
}

// PutTime puts time.Time variable to encoder.
func (e *Encoder) PutTime(t time.Time) (err error) {
	e.encodeKey()
	_, err = e.encodeTime(t.UTC())
	return
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
