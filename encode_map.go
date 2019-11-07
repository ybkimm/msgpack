package msgpack

func MarshalMap(o Map) ([]byte, error) {
	return NewEncoder(nil).encodeMap(o)
}

func (e *Encoder) PutMap(o Map) {
	e.encodeMap(o)
}

func (e *Encoder) PutMapKey(key string, o Map) {
	e.encodeString(key)
	e.encodeMap(o)
}

func (e *Encoder) encodeMap(o Map) ([]byte, error) {
	keysize := o.KeySize()
	e.grow(512)

	switch {
	case keysize <= fixmapMaxLen:
		e.writeByte(fixmapPrefix | byte(keysize))

	case keysize <= map16MaxLen:
		e.writeByte(Map16)
		e.writeUint16(uint16(keysize))

	default:
		e.writeByte(Map32)
		e.writeUint32(keysize)
	}

	o.MarshalMsgpackMap(e)
	return e.buf, e.err
}
