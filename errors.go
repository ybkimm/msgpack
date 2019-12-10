package msgpack

import (
	"encoding/json"
	"errors"
	"fmt"
	"reflect"
)

var (
	ErrTooLongString   = errors.New("msgpack: string is too long")
	ErrTooLongBinary   = errors.New("msgpack: binary is too long")
	ErrTooBigMap       = errors.New("msgpack: map is too big")
	ErrTooBigExtension = errors.New("msgpack: extension is too big")
	ErrUnexpectedEOF   = errors.New("msgpack: unexpected EOF")
	ErrNoReader        = errors.New("msgpack: reader is nil")
	ErrDecodeNil       = errors.New("msgpack: nil value was passed")
	ErrInvalidTime     = errors.New("msgpack: invalid time")
)

type ErrUnsupportedType struct {
	Type reflect.Type
}

func (e *ErrUnsupportedType) Error() string {
	return fmt.Sprintf("msgpack: unsupported type: %s", e.Type)
}

type ErrUnexpectedJSONToken struct {
	Token json.Token
}

func (e *ErrUnexpectedJSONToken) Error() string {
	return fmt.Sprintf("msgpack: unexpected JSON token %T: '%v'", e.Token, e.Token)
}

type ErrUnexpectedByte struct {
	Byte     byte
	Position int
	Stack    []byte
}

func (e *ErrUnexpectedByte) Error() string {
	return fmt.Sprintf("msgpack: unexpected byte 0x%02X at position %d, stack:\n%s", e.Byte, e.Position, e.Stack)
}

type ErrUnexpectedExtensionType struct {
	Type int8
}

func (e *ErrUnexpectedExtensionType) Error() string {
	return fmt.Sprintf("msgpack: unexpected extension type %d", e.Type)
}

type ErrUnregisteredExtension struct {
	Type int8
}

func (e *ErrUnregisteredExtension) Error() string {
	return fmt.Sprintf("msgpack: unregistered extension type %d", e.Type)
}

type ErrUnsupported struct {
	name string
}

func (e *ErrUnsupported) Error() string {
	return fmt.Sprintf("msgpack: %s is not supported", e.name)
}

type ErrUnknownKey struct {
	Key string
}

func (e *ErrUnknownKey) Error() string {
	return fmt.Sprintf("msgpack: unknown key: %s", e.Key)
}
