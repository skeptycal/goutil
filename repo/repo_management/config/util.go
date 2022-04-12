package config

const (
	Space      byte = ' '
	NL         byte = '\n'
	BreakItMax      = 50
	MaxItMax        = 50
	ModeDir         = 0755
	ModeFile        = 0644
)

type (
	Any interface{}

	GetSetter interface {
		Get(key Any) (value Any, err error)
		Set(key Any, value Any) error
	}

	FileSaver interface {

		// Load reads the file into memory
		ReadFile() error

		// Save saves the file from memory
		WriteFile() error
	}
)

// MaxIt returns the portion of the string that
// is at most MaxItMax characters long. If the
// string is shorter than MaxItMax, the entire
// string is returned.
func MaxIt(s string) string {
	if len(s) > MaxItMax {
		return s[:MaxItMax]
	}
	return s
}

// MaxItemLen returns the length of the longest string in the slice
func MaxItemLen(items []string) (max int) {
	// if len(items) == 0 {
	// 	return 0
	// }

	// if len(items) == 1 {
	// 	return len(items[0])
	// }

	max = -1
	for _, item := range items {
		if len(item) > max {
			max = len(item)
		}
	}

	return
}
