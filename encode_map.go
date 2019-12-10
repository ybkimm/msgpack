package msgpack

// MarshalMap returns map value as msgpack format.
func MarshalMap(o Map) ([]byte, error) {
	return NewEncoder(nil).encodeMap(o)
}

// PutMap puts map variable to encoder.
func (e *Encoder) PutMap(o Map) (err error) {
	e.encodeKey()
	_, err = e.encodeMap(o)
	return
}

func (e *Encoder) encodeMap(o Map) (data []byte, err error) {
	e.grow(512)

	fields := o.Fields()
	keysize := len(fields)

	lastSize := len(e.buf)

	for _, field := range fields {
		e.key = field

		err = o.MarshalMsgpackMap(e, field)
		if err != nil {
			return
		}

		if len(e.key) > 0 {
			keysize--
		}
	}

	switch {
	case keysize <= fixmapMaxLen:
		e.insertByte(fixmapPrefix|byte(keysize), lastSize)

	case keysize <= map16MaxLen:
		e.insertUint16(uint16(keysize), lastSize)
		e.insertByte(Map16, lastSize)

	default:
		e.insertUint32(uint32(keysize), lastSize)
		e.insertByte(Map32, lastSize)
	}

	return e.buf, e.err
}
