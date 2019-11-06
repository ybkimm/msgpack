package msgpack

func (e *Encoder) EncodeMap(o Map) error {
	e.putMap(o)
	return e.flush()
}

func (e *Encoder) PutMap(o Map) {
	e.putMap(o)
}

func (e *Encoder) PutMapKey(key string, o Map) {
	e.putString(key)
	e.putMap(o)
}

func (e *Encoder) putMap(o Map) {
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
}
