package ansi

const (
	SymCommand            string = `⌘`
	SymPlusMinus          string = `±` // Option Shift =
	SymDegree             string = `°` // Option Shift 8
	SymSquared            string = `²` // (no default macOS keymapping)
	SymCubed              string = `³` // (no default macOS keymapping)
	SymUpsideDownQuestion string = `¿` // Option Shift /
	SymCopyright          string = `©` // Option g
	SymRegistered         string = `®` // Option r
	SymTrademark          string = `™` // Option 2
	SymSummation          string = `∑` // Option w
	SymFrontTick          string = `´` // Option e
	SymDager              string = `†` // Option t
	SymEyes               string = `¨` // Option u
	SymHighCarrot         string = `ˆ` // Option i
	SymPhi                string = `ø` // Option o
	SymDiameter           string = `⌀` // (no default macOS keymapping)
	SymSlashO             string = `Ø`
	SymSmallSlashO        string = `ø`
	SymEmptySet           string = `∅`
)

var currency = map[string]string{
	`Yen`:        `¥`, // Option y
	`Currency`:   `¤`,
	`Pound`:      `£`, // Option 3
	`SymEuro`:    `€`, // Option Shift 2
	`SymBitcoin`: `₿`, // (no default macOS keymapping)
	`SymCents`:   `¢`, // Option 4
	`SymPeso`:    `₱`, // (no default macOS keymapping)
}
