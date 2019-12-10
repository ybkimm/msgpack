package msgpack

const (
	positiveFixnumMax = 127
	negativeFixnumMin = -1

	fixstrMaxLen = 0b00011111
	str8MaxLen   = 0xFF
	str16MaxLen  = 0xFFFF
	str32MaxLen  = 0xFFFFFFFF
	bin8MaxLen   = 0xFF
	bin16MaxLen  = 0xFFFF
	bin32MaxLen  = 0xFFFFFFFF
	fixarrMaxLen = 0b00001111
	arr16MaxLen  = 0xFFFF
	arr32MaxLen  = 0xFFFFFFFF
	fixmapMaxLen = 0b00001111
	map16MaxLen  = 0xFFFF
	map32MaxLen  = 0xFFFFFFFF
	ext8MaxLen   = 0xFF
	ext16MaxLen  = 0xFFFF
	ext32MaxLen  = 0xFFFFFFFF

	fixstrPrefix = 0b10100000
	fixarrPrefix = 0b10010000
	fixmapPrefix = 0b10000000

	Nil      = 0xC0
	False    = 0xC2
	True     = 0xC3
	Uint8    = 0xCC
	Uint16   = 0xCD
	Uint32   = 0xCE
	Uint64   = 0xCF
	Int8     = 0xD0
	Int16    = 0xD1
	Int32    = 0xD2
	Int64    = 0xD3
	Float32  = 0xCA
	Float64  = 0xCB
	String8  = 0xD9
	String16 = 0xDA
	String32 = 0xDB
	Binary8  = 0xC4
	Binary16 = 0xC5
	Binary32 = 0xC6
	Array16  = 0xDC
	Array32  = 0xDD
	Map16    = 0xDE
	Map32    = 0xDF
	Fixext1  = 0xD4
	Fixext2  = 0xD5
	Fixext4  = 0xD6
	Fixext8  = 0xD7
	Fixext16 = 0xD8
	Ext8     = 0xC7
	Ext16    = 0xC8
	Ext32    = 0xC9

	TimestampExt = -1
)
