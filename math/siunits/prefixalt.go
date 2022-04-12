package siunits

import (
	"fmt"
	"strings"
)

/// ********************* not needed for this

type (

	// alternate version for performance profiling
	prefixSI struct {
		locked   bool
		prefixes map[int]string
	}
)

const siSEP = ":"

func parseSymbol(v string) []string {
	return strings.Split(v, siSEP)
}

func (si *prefixSI) Name(key int) Any {
	if v, err := si.Get(key); err == nil {
		return strings.Split(v.(string), siSEP)[0]
	}
	return nil
}

func (si *prefixSI) Symbol(key int) Any {
	if v, err := si.Get(key); err == nil {
		return strings.Split(v.(string), siSEP)[1]
	}
	return nil
}

func (si *prefixSI) Get(key Any) (Any, error) {
	if err := si.keyGuard(key); err != nil {
		return "", err
	}
	if v, ok := si.prefixes[key.(int)]; ok {
		return v, nil
	}
	return "", fmt.Errorf("key not found: %v", key)
}

func (si *prefixSI) Set(key, value Any) error {
	if err := si.guard(key, value); err != nil {
		return err
	}

	if _, ok := si.prefixes[key.(int)]; ok {
		if si.locked {
			return fmt.Errorf("duplicate key not allowed (data locked): %v", key)
		}
	}

	si.prefixes[key.(int)] = value.(string)
	return nil
}

func (si *prefixSI) Add(base int, name, symbol string) {}

func (si *prefixSI) keyGuard(key Any) error {
	switch key.(type) {
	case int:
		return nil
	default:
		return fmt.Errorf("wrong key type: %v(%T)", key, key)
	}
}

func (si *prefixSI) valueGuard(key Any) error {
	switch key.(type) {
	case string:
		return nil
	default:
		return fmt.Errorf("wrong value type: %v(%T)", key, key)
	}
}

func (si *prefixSI) guard(key, value Any) error {

	if err := si.keyGuard(key); err == nil {
		if _, ok := si.prefixes[key.(int)]; ok {
			return nil
		}
		return fmt.Errorf("key not found: %v", key)
	}
	err := si.keyGuard(key)
	if err != nil {
		return err
	}
	return si.valueGuard(value)
}
