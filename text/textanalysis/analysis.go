package textanalysis

type (
	Any interface{}

	Word interface {
		Node
	}

	// Word is a temporary and incomplete representation of a
	// word in a text source. It implements the Noder interface
	// and may be used as the Data structure of a Node.
	//
	// It is designed to be regularly modified to test
	// different statistical and anecdotal theories.
	//
	// Do not rely on a consistant interface or functionality.
	word struct {
		count      int
		synonyms   []string
		definition string
		wikipedia  string
	}

	// Words is a map of words in a text source. They keys
	// are the actual word as presented in the text. The values
	// are the corresponding structure of properties.
	Words map[string]Word
)
