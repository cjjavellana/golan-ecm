package timeutils

import "time"

// Undefined used when a nil is not possible. e.g. When a struct has a non-pointer time.Time field and must be
// set to value that is not determined yet.
//
// Returns January 1st 1900 00:00:00
func Undefined() time.Time {
	return time.Date(1900, time.January, 1, 0, 0, 0, 0, time.UTC)
}
