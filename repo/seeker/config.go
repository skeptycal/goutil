package seeker

import (
	"reflect"
)

const (
	defaultServer = "localhost"
)

type (
	Any    interface{}
	AnyMap = map[Any]Any

	AnyType    reflect.Type
	AnyTypeMap = map[AnyType]AnyType

	Settings struct {
		portString string
		handlers   HandlerMap
	}
)

type StringMap struct {
	name         string
	writeProtect bool
	keyType      Any
	m            map[string]Any
}

type HandlerMap struct {
	name         string
	writeProtect bool
	keyType      Any
	m            map[string]Handler
}

// func (m HandlerMap) Get(key Any) Any {

// 	if s, ok := key.(string); ok {
// 		if value, ok := m[s]; ok {
// 			return value
// 		}

// 		// "normal" return if string type key is used but value is not present
// 		return nil
// 	}

// 	// "fallback" return if key type is not "string"
// 	// could log error or panic here to expose any errors
// 	return nil
// }
