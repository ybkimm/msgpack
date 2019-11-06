package msgpack

import "time"

func (e *Encoder) EncodeTime(t time.Time) error {
	e.putTime(t.UTC())
	return e.flush()
}

func (e *Encoder) PutTime(t time.Time) {
	e.putTime(t.UTC())
}

func (e *Encoder) PutTimeKey(key string, t time.Time) {
	e.putString(key)
	e.putTime(t.UTC())
}

func (e *Encoder) putTime(t time.Time) {
	var (
		seconds     = uint64(t.Unix())
		nanoseconds = uint64(t.UnixNano()) - 1000000000*seconds
	)

	e.putExtension(&extTime{
		seconds:     seconds,
		nanoseconds: nanoseconds,
	})
}
