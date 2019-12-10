package msgpack

import (
	"encoding/binary"
	"time"
)

var _ Extension = (*extTime)(nil)

type extTime struct {
	seconds     uint64
	nanoseconds uint64
}

func (t *extTime) ExtensionType() int8 {
	return TimestampExt
}

func (t *extTime) MarshalMsgpackExtension() []byte {
	if (t.seconds>>32) == 0 && t.nanoseconds == 0 {
		return []byte{
			byte(t.seconds >> 24),
			byte(t.seconds >> 16),
			byte(t.seconds >> 8),
			byte(t.seconds),
		}
	} else if (t.seconds >> 34) == 0 {
		data := t.nanoseconds<<34 | t.seconds
		return []byte{
			byte(data >> 56),
			byte(data >> 48),
			byte(data >> 40),
			byte(data >> 32),
			byte(data >> 24),
			byte(data >> 16),
			byte(data >> 8),
			byte(data),
		}
	} else {
		return []byte{
			byte(t.nanoseconds >> 24),
			byte(t.nanoseconds >> 16),
			byte(t.nanoseconds >> 8),
			byte(t.nanoseconds),
			byte(t.seconds >> 56),
			byte(t.seconds >> 48),
			byte(t.seconds >> 40),
			byte(t.seconds >> 32),
			byte(t.seconds >> 24),
			byte(t.seconds >> 16),
			byte(t.seconds >> 8),
			byte(t.seconds),
		}
	}
}

func (t *extTime) UnmarshalMsgpackExtension(p []byte) error {
	switch len(p) {
	case 4:
		t.nanoseconds = 0
		t.seconds = uint64(binary.BigEndian.Uint32(p))

	case 8:
		data := binary.BigEndian.Uint64(p)
		t.nanoseconds = data >> 34
		t.seconds = data & 0x00000003FFFFFFF

	case 12:
		t.nanoseconds = uint64(binary.BigEndian.Uint32(p))
		t.seconds = binary.BigEndian.Uint64(p[4:])

	default:
		return ErrInvalidTime
	}

	return nil
}

func (t *extTime) Time() time.Time {
	return time.Unix(int64(t.seconds), int64(t.nanoseconds))
}
