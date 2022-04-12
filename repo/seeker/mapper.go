package seeker

import (
	"errors"
	"reflect"
)

const minMapperSize int = 8

type GetSetter interface {
	Get(key Any) (Any, error)
	Set(key, value Any) error
}

type Mapper interface {
	GetSetter
	KeyType() reflect.Type
	Len() int
}

type Datapoint interface {
	GetSetter
}

// LineMapper maps items to sequentially numbered
// integer keys
type LineMapper interface {
	Mapper
	Counter() int
}

func NewMapper(name string, size int, writeProtect bool, typeChecking bool, keyType reflect.Type) (Mapper, error) {
	if size < minMapperSize {
		size = minMapperSize
	}

	// t := reflect.TypeOf(sampleKey)

	m := &mapper{
		name:         name,
		writeProtect: writeProtect,
		keyType:      keyType,
		keyTypeCheck: typeChecking,
		m:            make(AnyMap, size),
	}
	return m, nil
}

type mapper struct {
	name         string
	writeProtect bool

	// todo - is this worth implmenting?
	keyType      reflect.Type
	keyTypeCheck bool

	m map[Any]Any
}

//--------> Public Methods

// Get returns the value for the given key. If the
// key is not found, an error is returned.
//
// If the dynamic key type is incorrect for the mapper,
// and error is returned.
func (mpr *mapper) Get(key Any) (Any, error) {

	if !mpr.keyTypeOK(key) {
		return nil, ErrWrongKeyType
	}

	if value, ok := mpr.m[key]; ok {
		return value, nil
	}

	return nil, ErrKeyNotFound
}

// Set sets the value for the given key.
//
// If the mapper is write protected, only unset
// key values may be updated. If the key exists,
// an error is returned.
//
// If the dynamic key type is incorrect for the mapper,
// and error is returned.
func (mpr *mapper) Set(key, value Any) error {
	if !mpr.keyTypeOK(key) {
		return ErrWrongKeyType
	}

	if mpr.writeProtect {
		if _, ok := mpr.m[key]; ok {
			return ErrSetNotAllowed
		}
	}

	mpr.m[key] = value
	return nil
}

func (mpr *mapper) Len() int {
	return len(mpr.m)
}

//--------> Private Methods

// keyTypeOkOn returns true if the dynamic key is
// of the correct type for the mapper
func (mpr *mapper) keyTypeOK(key Any) bool {

	// todo - have turn on/off with a setting
	// or assign a function to a variable
	if mpr.keyTypeCheck {
		return mpr.keyTypeOkOn(key)

	}
	return mpr.keyTypeOkOff(key)
}

// keyTypeOkOn is the function used if dynamic key
// type checking is turned off.
func (mpr *mapper) keyTypeOkOff(key Any) bool { return true }

// keyTypeOkOn is the function used if dynamic key
// type checking is turned on.
func (mpr *mapper) keyTypeOkOn(key Any) bool {
	if reflect.TypeOf(key) != mpr.keyType {
		return false
	}
	return true
}

//--------> Errors

var (
	ErrWrongKeyType  = errors.New("mapper: wrong key type")
	ErrKeyNotFound   = errors.New("mapper: key not found")
	ErrSetNotAllowed = errors.New("mapper: command Set() not allowed when map is write protected")
)
