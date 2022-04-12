package config

type (
	// SyncMapper implements the public interface for sync.Map
	SyncMapper interface {

		// Load returns the value stored in the map for a key, or nil if no
		// value is present.
		// The ok result indicates whether value was found in the map.
		Load(key interface{}) (value interface{}, ok bool)

		// Store sets the value for a key.
		Store(key, value interface{})

		// LoadOrStore returns the existing value for the key if present.
		// Otherwise, it stores and returns the given value.
		// The loaded result is true if the value was loaded, false if stored.
		LoadOrStore(key, value interface{}) (actual interface{}, loaded bool)

		// LoadAndDelete deletes the value for a key, returning the previous value if any.
		// The loaded result reports whether the key was present.
		LoadAndDelete(key interface{}) (value interface{}, loaded bool)

		// Delete deletes the value for a key.
		Delete(key interface{})

		// Range calls f sequentially for each key and value present in the map.
		// If f returns false, range stops the iteration.
		//
		// Range does not necessarily correspond to any consistent snapshot of the Map's
		// contents: no key will be visited more than once, but if the value for any key
		// is stored or deleted concurrently, Range may reflect any mapping for that key
		// from any point during the Range call.
		//
		// Range may be O(N) with the number of elements in the map even if f returns
		// false after a constant number of calls.
		Range(f func(key, value interface{}) bool)
	}
)
