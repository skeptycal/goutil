package pooler

import "errors"

type (
	GetSetter interface {
		Get(key Any) (Any, error)
		Set(key Any, value Any) error
	}

	Params interface {
		GetSetter
	}

	params struct {
		m      map[string]Any
		locked bool
	}
)

func NewParams(locked bool, m map[string]Any) Params {

	return params{
		m:      m,
		locked: locked,
	}

}

func (p params) Set(key Any, value Any) error {
	if p.locked {
		return errValueLocked
	}

	if _, err := p.Get(key); err != nil {
		if err != errKeyNotFound {
			return err
		}
	}

	p.m[key.(string)] = value
	return nil

}

func (p params) Get(key Any) (Any, error) {
	k, ok := key.(string)
	if !ok {
		return nil, errNotString
	}

	if k == "" {
		return nil, errNoKeyProvided
	}

	if v, ok := p.m[k]; ok {
		return v, nil
	}
	return nil, errKeyNotFound
}

var (
	errNotString     = errors.New("cannot create string from key")
	errNoKeyProvided = errors.New("no key provided")
	errKeyNotFound   = errors.New("key not found")
	errValueLocked   = errors.New("parameter values are locked")
)
