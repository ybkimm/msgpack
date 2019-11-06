package msgpack

// EncodeString encodes a string to msgpack format
func (e *Encoder) EncodeString(v string) error {
	e.putString(v)
	return e.flush()
}

func (e *Encoder) PutString(v string) {
	e.putString(v)
}

func (e *Encoder) PutStringKey(key string, v string) {
	e.putString(key)
	e.putString(v)
}

func (e *Encoder) putString(v string) {
	strlen := len(v)
	switch {
	case strlen <= fixstrMaxLen:
		e.grow(strlen + 1)
		e.writeByte(fixstrPrefix | byte(strlen))

	case strlen <= str8MaxLen:
		e.grow(strlen + 2)
		e.writeBytes([]byte{String8, byte(strlen)})

	case strlen <= str16MaxLen:
		e.grow(strlen + 3)
		e.writeByte(String16)
		e.writeUint16(uint16(strlen))

	case strlen <= str32MaxLen:
		e.grow(strlen + 5)
		e.writeByte(String32)
		e.writeUint32(uint32(strlen))

	default:
		panic(ErrTooLongString)
	}

	e.writeString(v)
}
