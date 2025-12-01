// Copyright © 2025 Ping Identity Corporation

package types

import (
	"encoding/json"
	"fmt"
	"time"
)

// UnixTime wraps time.Time to support JSON marshaling and unmarshaling using Unix timestamp integers.
// Times are always stored and returned in UTC. The zero value for UnixTime is equivalent to
// the Unix epoch (January 1, 1970 00:00:00 UTC).
//
// When unmarshaling, UnixTime accepts:
//   - Integer Unix timestamps in seconds (e.g., 1609459200)
//   - Integer Unix timestamps in milliseconds (e.g., 1609459200000)
//   - Automatically detects milliseconds vs seconds based on magnitude
//   - null values (results in zero time)
//
// When marshaling, UnixTime produces:
//   - Integer Unix timestamps in seconds for non-zero times
//   - null for zero times
//
// Example usage:
//
//	type Event struct {
//	    OccurredAt UnixTime `json:"occurred_at"`
//	}
//
//	// Unmarshaling from JSON
//	jsonData := []byte(`{"occurred_at": 1609459200}`)
//	var event Event
//	json.Unmarshal(jsonData, &event)
//
//	// Marshaling to JSON
//	data, _ := json.Marshal(event)
type UnixTime struct {
	time.Time
}

// UnmarshalJSON implements the json.Unmarshaler interface.
// It accepts Unix timestamp integers in seconds or milliseconds and converts them to UTC time.
// Automatically detects milliseconds (values > 1e12) vs seconds.
// Null values are accepted and result in a zero time value.
// Negative timestamps are supported for dates before the Unix epoch.
func (u *UnixTime) UnmarshalJSON(b []byte) error {
	// Handle null values
	if string(b) == "null" {
		u.Time = time.Time{}
		return nil
	}

	var timestamp int64
	if err := json.Unmarshal(b, &timestamp); err != nil {
		return fmt.Errorf("failed to unmarshal unix timestamp: %w", err)
	}

	// Auto-detect milliseconds vs seconds
	// Timestamps > 1e12 (1 trillion) are likely in milliseconds
	// This threshold corresponds to September 2001 in seconds or January 1970 in milliseconds
	if timestamp > 1e12 || timestamp < -1e12 {
		// Milliseconds since epoch
		u.Time = time.Unix(0, timestamp*int64(time.Millisecond)).UTC()
	} else {
		// Seconds since epoch
		u.Time = time.Unix(timestamp, 0).UTC()
	}

	return nil
}

// MarshalJSON implements the json.Marshaler interface.
// It converts the time to a Unix timestamp integer (seconds since epoch).
// Zero times are marshaled as null.
func (u UnixTime) MarshalJSON() ([]byte, error) {
	// Marshal zero times as null for consistency
	if u.Time.IsZero() {
		return []byte("null"), nil
	}

	return json.Marshal(u.Unix())
}

// String returns the time formatted in RFC3339 format with UTC timezone.
// This provides a human-readable representation for logging and debugging.
func (u UnixTime) String() string {
	if u.Time.IsZero() {
		return "<zero>"
	}
	return u.Time.UTC().Format(time.RFC3339)
}
