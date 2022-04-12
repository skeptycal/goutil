package defaults

type Setting interface {
	Booler
}

func NewSetting(key string, value Any) Setting { return AnyBooler(value) }
