package config

type (

	// configFlag represents a single command line flag
	configFlag struct{}

	// configFlagSet represents a set of command line flags
	configFlagSet struct {
		id    int
		name  string
		flags map[string]configFlag
	}

	Flagger interface {
		GetSetter
	}
)
