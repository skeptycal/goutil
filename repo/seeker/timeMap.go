package seeker

import "time"

type (

	// A LocalTime represents an instant in time with
	// nanosecond precision. It is a modified version
	// of the standard library time package Time type
	// that removes timezone information for performance
	// reasons.
	//
	// The goal is to create a time type that specifies
	// values that can be used as unique keys in a
	// database or map. From the standard library:
	// "Time values should not be used as map or database keys without first guaranteeing that the identical Location has been set for all values"
	//
	// Reference: go 1.17.5
	//
	// Programs using times should typically store and pass them as values,
	// not pointers. That is, time variables and struct fields should be of
	// type time.Time, not *time.Time.
	//
	// A Time value can be used by multiple goroutines simultaneously except
	// that the methods GobDecode, UnmarshalBinary, UnmarshalJSON and
	// UnmarshalText are not concurrency-safe.
	//
	// Time instants can be compared using the Before, After, and Equal methods.
	// The Sub method subtracts two instants, producing a Duration.
	// The Add method adds a Time and a Duration, producing a Time.
	//
	// The zero value of type Time is January 1, year 1, 00:00:00.000000000 UTC.
	// As this time is unlikely to come up in practice, the IsZero method gives
	// a simple way of detecting a time that has not been initialized explicitly.
	//
	// In addition to the required “wall clock” reading, a Time may contain an optional
	// reading of the current process's monotonic clock, to provide additional precision
	// for comparison or subtraction.
	// See the “Monotonic Clocks” section in the package documentation for details.
	//
	// Note that the Go == operator compares not just the time instant but also the
	// Location and the monotonic clock reading. Therefore, Time values should not
	// be used as map or database keys without first guaranteeing that the
	// identical Location has been set for all values, which can be achieved
	// through use of the UTC or Local method, and that the monotonic clock reading
	// has been stripped by setting t = t.Round(0). In general, prefer t.Equal(u)
	// to t == u, since t.Equal uses the most accurate comparison available and
	// correctly handles the case when only one of its arguments has a monotonic
	// clock reading.
	//
	LocalTime struct {
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
	// TimeDataPoint represents data related to
	// a specific time interval.
	TimeDataPoint interface{}

	// timeMap is an ordered map of time series values
	TimeMap map[int64]TimeDataPoint
)

var tt time.Time

func (t TimeMap) Get(time int64) (TimeDataPoint, bool) {
	if td, ok := t[time]; ok {
		return td, true
	}

	return nil, false
}
