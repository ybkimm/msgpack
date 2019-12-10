package msgpack

// Extension is an interface for application-specific type.
type Extension interface {
	ExtensionType() int8
	MarshalMsgpackExtension() []byte
	UnmarshalMsgpackExtension(p []byte) error
}

// NullableExtension is an extension that can be nil.
type NullableExtension interface {
	Nullable
	Extension
}

type nullableExtension struct {
	Extension
}

func (ne *nullableExtension) UnmarshalMsgpackNull() error {
	return ne.UnmarshalMsgpackExtension(nil)
}

func unwrapNullableExtension(e Extension) Extension {
	ne, ok := e.(*nullableExtension)
	if !ok {
		return e
	}
	return ne.Extension
}
