package msgpack

import "time"

type Array interface {
	MarshalMsgpackArray(e *Encoder, i int) error
	UnmarshalMsgpackArray(d *Decoder, l int) error
	Length() uint32
}

type NullableArray interface {
	Nullable
	Array
}

type BoolArray []bool

func (a *BoolArray) MarshalMsgpackArray(e *Encoder, i int) error {
	return e.PutBool((*a)[i])
}

func (a *BoolArray) UnmarshalMsgpackArray(d *Decoder, l int) error {
	var err error
	*a = make([]bool, l, l)
	for i := 0; i < l; i++ {
		err = d.DecodeBool(&(*a)[i])
		if err != nil {
			return err
		}
	}
	return nil
}

func (a *BoolArray) Length() uint32 {
	return uint32(len(*a))
}

type IntArray []int

func (a *IntArray) MarshalMsgpackArray(e *Encoder, i int) error {
	return e.PutInt((*a)[i])
}

func (a *IntArray) UnmarshalMsgpackArray(d *Decoder, l int) error {
	var err error
	*a = make([]int, l, l)
	for i := 0; i < l; i++ {
		err = d.DecodeInt(&(*a)[i])
		if err != nil {
			return err
		}
	}
	return nil
}

func (a *IntArray) Length() uint32 {
	return uint32(len(*a))
}

type Int8Array []int8

func (a *Int8Array) MarshalMsgpackArray(e *Encoder, i int) error {
	return e.PutInt8((*a)[i])
}

func (a *Int8Array) UnmarshalMsgpackArray(d *Decoder, l int) error {
	var err error
	*a = make([]int8, l, l)
	for i := 0; i < l; i++ {
		err = d.DecodeInt8(&(*a)[i])
		if err != nil {
			return err
		}
	}
	return nil
}

func (a *Int8Array) Length() uint32 {
	return uint32(len(*a))
}

type Int16Array []int16

func (a *Int16Array) MarshalMsgpackArray(e *Encoder, i int) error {
	return e.PutInt16((*a)[i])
}

func (a *Int16Array) UnmarshalMsgpackArray(d *Decoder, l int) error {
	var err error
	*a = make([]int16, l, l)
	for i := 0; i < l; i++ {
		err = d.DecodeInt16(&(*a)[i])
		if err != nil {
			return err
		}
	}
	return nil
}

func (a *Int16Array) Length() uint32 {
	return uint32(len(*a))
}

type Int32Array []int32

func (a *Int32Array) MarshalMsgpackArray(e *Encoder, i int) error {
	return e.PutInt32((*a)[i])
}

func (a *Int32Array) UnmarshalMsgpackArray(d *Decoder, l int) error {
	var err error
	*a = make([]int32, l, l)
	for i := 0; i < l; i++ {
		err = d.DecodeInt32(&(*a)[i])
		if err != nil {
			return err
		}
	}
	return nil
}

func (a *Int32Array) Length() uint32 {
	return uint32(len(*a))
}

type Int64Array []int64

func (a *Int64Array) MarshalMsgpackArray(e *Encoder, i int) error {
	return e.PutInt64((*a)[i])
}

func (a *Int64Array) UnmarshalMsgpackArray(d *Decoder, l int) error {
	var err error
	*a = make([]int64, l, l)
	for i := 0; i < l; i++ {
		err = d.DecodeInt64(&(*a)[i])
		if err != nil {
			return err
		}
	}
	return nil
}

func (a *Int64Array) Length() uint32 {
	return uint32(len(*a))
}

type UintArray []uint

func (a *UintArray) MarshalMsgpackArray(e *Encoder, i int) error {
	return e.PutUint((*a)[i])
}

func (a *UintArray) UnmarshalMsgpackArray(d *Decoder, l int) error {
	var err error
	*a = make([]uint, l, l)
	for i := 0; i < l; i++ {
		err = d.DecodeUint(&(*a)[i])
		if err != nil {
			return err
		}
	}
	return nil
}

func (a *UintArray) Length() uint32 {
	return uint32(len(*a))
}

type Uint16Array []uint16

func (a *Uint16Array) MarshalMsgpackArray(e *Encoder, i int) error {
	return e.PutUint16((*a)[i])
}

func (a *Uint16Array) UnmarshalMsgpackArray(d *Decoder, l int) error {
	var err error
	*a = make([]uint16, l, l)
	for i := 0; i < l; i++ {
		err = d.DecodeUint16(&(*a)[i])
		if err != nil {
			return err
		}
	}
	return nil
}

func (a *Uint16Array) Length() uint32 {
	return uint32(len(*a))
}

type Uint32Array []uint32

func (a *Uint32Array) MarshalMsgpackArray(e *Encoder, i int) error {
	return e.PutUint32((*a)[i])
}

func (a *Uint32Array) UnmarshalMsgpackArray(d *Decoder, l int) error {
	var err error
	*a = make([]uint32, l, l)
	for i := 0; i < l; i++ {
		err = d.DecodeUint32(&(*a)[i])
		if err != nil {
			return err
		}
	}
	return nil
}

func (a *Uint32Array) Length() uint32 {
	return uint32(len(*a))
}

type Uint64Array []uint64

func (a *Uint64Array) MarshalMsgpackArray(e *Encoder, i int) error {
	return e.PutUint64((*a)[i])
}

func (a *Uint64Array) UnmarshalMsgpackArray(d *Decoder, l int) error {
	var err error
	*a = make([]uint64, l, l)
	for i := 0; i < l; i++ {
		err = d.DecodeUint64(&(*a)[i])
		if err != nil {
			return err
		}
	}
	return nil
}

func (a *Uint64Array) Length() uint32 {
	return uint32(len(*a))
}

type Float32Array []float32

func (a *Float32Array) MarshalMsgpackArray(e *Encoder, i int) error {
	return e.PutFloat32((*a)[i])
}

func (a *Float32Array) UnmarshalMsgpackArray(d *Decoder, l int) error {
	var err error
	*a = make([]float32, l, l)
	for i := 0; i < l; i++ {
		err = d.DecodeFloat32(&(*a)[i])
		if err != nil {
			return err
		}
	}
	return nil
}

func (a *Float32Array) Length() uint32 {
	return uint32(len(*a))
}

type Float64Array []float64

func (a *Float64Array) MarshalMsgpackArray(e *Encoder, i int) error {
	return e.PutFloat64((*a)[i])
}

func (a *Float64Array) UnmarshalMsgpackArray(d *Decoder, l int) error {
	var err error
	*a = make([]float64, l, l)
	for i := 0; i < l; i++ {
		err = d.DecodeFloat64(&(*a)[i])
		if err != nil {
			return err
		}
	}
	return nil
}

func (a *Float64Array) Length() uint32 {
	return uint32(len(*a))
}

type StringArray []string

func (a *StringArray) MarshalMsgpackArray(e *Encoder, i int) error {
	return e.PutString((*a)[i])
}

func (a *StringArray) UnmarshalMsgpackArray(d *Decoder, l int) error {
	var err error
	*a = make([]string, l, l)
	for i := 0; i < l; i++ {
		err = d.DecodeString(&(*a)[i])
		if err != nil {
			return err
		}
	}
	return nil
}

func (a *StringArray) Length() uint32 {
	return uint32(len(*a))
}

type BinaryArray [][]byte

func (a *BinaryArray) MarshalMsgpackArray(e *Encoder, i int) error {
	return e.PutBinary((*a)[i])
}

func (a *BinaryArray) UnmarshalMsgpackArray(d *Decoder, l int) error {
	var err error
	*a = make([][]byte, l, l)
	for i := 0; i < l; i++ {
		err = d.DecodeBinary(&(*a)[i])
		if err != nil {
			return err
		}
	}
	return nil
}

func (a *BinaryArray) Length() uint32 {
	return uint32(len(*a))
}

type TimeArray []time.Time

func (a *TimeArray) MarshalMsgpackArray(e *Encoder, i int) error {
	return e.PutTime((*a)[i])
}

func (a *TimeArray) UnmarshalMsgpackArray(d *Decoder, l int) error {
	var err error
	*a = make([]time.Time, l, l)
	for i := 0; i < l; i++ {
		err = d.DecodeTime(&(*a)[i])
		if err != nil {
			return err
		}
	}
	return nil
}

func (a *TimeArray) Length() uint32 {
	return uint32(len(*a))
}
