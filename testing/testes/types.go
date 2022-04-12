package testes

import "

var (
	ValueOf = types.ValueOf
)

type (
	Any = types.Any

	AnyValue = types.AnyValue

	GetSetter interface {
		Get(key Any) (Any, error)
		Set(key Any, value Any) error
	}
)
