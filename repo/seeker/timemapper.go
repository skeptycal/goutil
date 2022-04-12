package seeker

import (
	"reflect"
	"time"
)

type TimeData interface {
	Timestamp() time.Time
	Datapoint() []Datapoint
}

type TimeSeries interface {
	Start() time.Time
	Stop() time.Time
	Data() []TimeData
	First() TimeData
	Next() TimeData
}

// TimeMapper is based on Mapper
//
//	 type Mapper interface {
//	 	GetSetter
//	 	KeyType() reflect.Type
//	 	Len() int
//	 }
type TimeMapper interface {
	Mapper
	TimeSlice(start, end time.Time) []TimeSeries
}

// timeMapper is a data structure based on mapper:
//
// 	type mapper struct {
// 		name         string
// 		writeProtect bool
// 		keyType      reflect.Type
// 		keyTypeCheck bool
// 		m map[Any]Any
// 	}
//
// keyType is always time.Time (no checks needed)
type timeMapper struct {
	mapper
}

var timeType = reflect.TypeOf(time.Now())

func (mpr *timeMapper) KeyType() reflect.Type {
	return timeType
}

// Get returns the value for the given key. If the
// key is not found, an error is returned.
//
// If the dynamic key type is incorrect for the mapper,
// and error is returned.
func (mpr *timeMapper) Get(key Any) (Any, error) {

	if key.(timeType)

	if !mpr.keyTypeOK(key) {
		return nil, ErrWrongKeyType
	}

	if value, ok := mpr.m[key]; ok {
		return value, nil
	}

	return nil, ErrKeyNotFound
}
