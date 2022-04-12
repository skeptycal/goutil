package seeker

// alpha and omega are the beginning and end of time for zone
// transitions.
//
// (much borrowed from Go standard library time/zoneinfo.go 1.17.5)
const (
	alpha = -1 << 63  // math.MinInt64
	omega = 1<<63 - 1 // math.MaxInt64
)

type (
	Time struct {
		// wall and ext encode the wall time seconds, wall time nanoseconds,
		// and optional monotonic clock reading in nanoseconds.
		//
		// From high to low bit position, wall encodes a 1-bit flag (hasMonotonic),
		// a 33-bit seconds field, and a 30-bit wall time nanoseconds field.
		// The nanoseconds field is in the range [0, 999999999].
		// If the hasMonotonic bit is 0, then the 33-bit field must be zero
		// and the full signed 64-bit wall seconds since Jan 1 year 1 is stored in ext.
		// If the hasMonotonic bit is 1, then the 33-bit field holds a 33-bit
		// unsigned wall seconds since Jan 1 year 1885, and ext holds a
		// signed 64-bit monotonic clock reading, nanoseconds since process start.
		wall uint64
		ext  int64

		// loc specifies the Location that should be used to
		// determine the minute, hour, month, day, and year
		// that correspond to this Time.
		// The nil location means UTC.
		// All UTC times are represented with loc==nil, never loc==&utcLoc.
		loc *Location
	}

	// A Location maps time instants to the zone in use at that time.
	// Typically, the Location represents the collection of time offsets
	// in use in a geographical area. For many Locations the time offset varies
	// depending on whether daylight savings time is in use at the time instant.
	Location struct {
		name string
		zone []zone
		tx   []zoneTrans

		// The tzdata information can be followed by a string that describes
		// how to handle DST transitions not recorded in zoneTrans.
		// The format is the TZ environment variable without a colon; see
		// https://pubs.opengroup.org/onlinepubs/9699919799/basedefs/V1_chap08.html.
		// Example string, for America/Los_Angeles: PST8PDT,M3.2.0,M11.1.0
		extend string

		// Most lookups will be for the current time.
		// To avoid the binary search through tx, keep a
		// static one-element cache that gives the correct
		// zone for the time when the Location was created.
		// if cacheStart <= t < cacheEnd,
		// lookup can return cacheZone.
		// The units for cacheStart and cacheEnd are seconds
		// since January 1, 1970 UTC, to match the argument
		// to lookup.
		cacheStart int64
		cacheEnd   int64
		cacheZone  *zone
	}

	// A zone represents a single time zone such as CET.
	zone struct {
		name   string // abbreviated name, "CET"
		offset int    // seconds east of UTC
		isDST  bool   // is this zone Daylight Savings Time?
	}

	// A zoneTrans represents a single time zone transition.
	zoneTrans struct {
		when         int64 // transition time, in seconds since 1970 GMT
		index        uint8 // the index of the zone that goes into effect at that time
		isstd, isutc bool  // ignored - no idea what these mean
	}
)
