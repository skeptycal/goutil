package greek

import (
	"math"
	"math/rand"
	"strings"
	"time"

	"github.com/skeptycal/goutil/types"
)

func init() {
	rand.Seed(int64(time.Now().Nanosecond()))
}

func RandomGreekString(n int) string {
	sb := strings.Builder{}
	defer sb.Reset()

	keys := Greek.Keys()
	listLen := len(keys) - 1

	for i := 0; i < n; i++ {
		pos := rand.Intn(listLen)
		if rand.Intn(100) < 50 {
			sb.WriteRune(Greek[keys[pos]].lower)
		} else {
			sb.WriteRune(Greek[keys[pos]].upper)

		}
	}

	return sb.String()
}

func RandomGreek(length, minGreekPct, maxGreekPct int) string {
	rng := maxGreekPct - minGreekPct
	pct := rand.Intn(rng) + minGreekPct
	numGreek := int(math.Floor(float64(pct) / 100 * float64(length)))
	// get regular text

	s := types.RandomString(length - numGreek)

	// get Greek text
	s += RandomGreekString(numGreek)

	b := []rune(s)

	rand.Shuffle(len(b), func(i, j int) { b[i], b[j] = b[j], b[i] })
	return string(b)
}

type charMap map[string]struct {
	English rune
	upper   rune
	lower   rune
}

func (c charMap) toLower(s string) rune {
	if v, ok := c[s]; ok {
		return v.lower
	}
	return types.ReplacementChar
}
func (c charMap) toUpper(s string) rune {
	if v, ok := c[s]; ok {
		return v.upper
	}
	return types.ReplacementChar
}

func (c charMap) ToLower(s string) string {

	/// TODO: this will be a slow algorithm ...
	for _, v := range Greek {
		if strings.Contains(s, string(v.upper)) {
			s = strings.ReplaceAll(s, string(v.upper), string(v.lower))
		}
	}
	return s
}
func (c charMap) ToUpper(s string) string {

	/// TODO: this will be a slow algorithm ...
	for _, v := range Greek {
		if strings.Contains(s, string(v.lower)) {
			s = strings.ReplaceAll(s, string(v.lower), string(v.upper))
		}
	}
	return s
}

func (c charMap) Len() int { return len(c) }

func (c charMap) Keys() []string {
	keys := make([]string, 0, c.Len())
	for k := range c {
		keys = append(keys, k)
	}
	return keys
}

var Greek = charMap{
	// Name, Uppercase, Lowercase
	"Alpha":   {'A', 'Α', 'α'},
	"Beta":    {'B', 'Β', 'β'},
	"Gamma":   {'G', 'Γ', 'γ'},
	"Delta":   {'D', 'Δ', 'δ'},
	"Epsilon": {'E', 'Ε', 'ε'},
	"Zeta":    {'Z', 'Ζ', 'ζ'},
	"Eta":     {'H', 'Η', 'η'},
	"Theta":   {'T', 'Θ', 'θ'},
	"Iota":    {'I', 'Ι', 'ι'},
	"Kappa":   {'K', 'Κ', 'κ'},
	"Lambda":  {'L', 'Λ', 'λ'},
	"Mu":      {'M', 'Μ', 'μ'},
	"Nu":      {'N', 'Ν', 'ν'},
	"Xi":      {'X', 'Ξ', 'ξ'},
	"Omicron": {'O', 'Ο', 'ο'},
	"Pi":      {'P', 'Π', 'π'},
	"Rho":     {'R', 'Ρ', 'ρ'},
	"Sigma":   {'S', 'Σ', 'σ'},
	"Tau":     {'t', 'Τ', 'τ'},
	"Upsilon": {'U', 'Υ', 'υ'},
	"Phi":     {'p', 'Φ', 'φ'},
	"Chi":     {'C', 'Χ', 'χ'},
	"Psi":     {'s', 'Ψ', 'ψ'},
	"Omega":   {'W', 'Ω', 'ω'},
}
