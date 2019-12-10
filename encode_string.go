package msgpack

// MarshalString returns string value as msgpack format.
func MarshalString(v string) ([]byte, error) {
	return NewEncoder(nil).encodeString(v)
}

// PutString puts string variable to encoder.
func (e *Encoder) PutString(v string) (err error) {
	e.encodeKey()
	_, err = e.encodeString(v)
	return
}

func (e *Encoder) encodeString(v string) ([]byte, error) {
	strlen := len(v)
	switch {
	case strlen <= fixstrMaxLen:
		e.grow(strlen + 1)
		e.writeByte(fixstrPrefix | byte(strlen))
		e.writeString(v)

	case strlen <= str8MaxLen:
		e.grow(strlen + 2)
		e.writeBytes([]byte{String8, byte(strlen)})
		e.writeString(v)

	case strlen <= str16MaxLen:
		e.grow(strlen + 3)
		e.writeByte(String16)
		e.writeUint16(uint16(strlen))
		e.writeString(v)

	case strlen <= str32MaxLen:
		e.grow(strlen + 5)
		e.writeByte(String32)
		e.writeUint32(uint32(strlen))
		e.writeString(v)

	default:
		e.err = ErrTooLongString
	}

	return e.buf, e.err
}
