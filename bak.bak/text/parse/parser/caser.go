package parser

type Caser interface {
	String() string
	ChangeCase(c Cases)
}

type caser struct {
	in  string
	out string
	c   Cases
}

// String returns a human-readable representation
// of the string with the corresponding case.
func (ca *caser) String() string {
	if ca.out == "" {
		ca.process()
	}
	return ca.out
}

// ChangeCase changes the output case of the string.
func (ca *caser) ChangeCase(c Cases) {
	ca.c = c
	ca.process()
}

func (ca *caser) process() {
	ca.out = process(ca.c, ca.in)
}
