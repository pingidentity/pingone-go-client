// Copyright © 2025 Ping Identity Corporation

package types

import (
	"encoding/json"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestUnixTime_UnmarshalJSON(t *testing.T) {
	tests := []struct {
		name        string
		input       string
		expected    time.Time
		expectError bool
	}{
		{
			name:     "seconds timestamp",
			input:    "1609459200",
			expected: time.Unix(1609459200, 0).UTC(),
		},
		{
			name:     "milliseconds timestamp",
			input:    "1764547200000",
			expected: time.Unix(0, 1764547200000*int64(time.Millisecond)).UTC(),
		},
		{
			name:     "zero timestamp",
			input:    "0",
			expected: time.Unix(0, 0).UTC(),
		},
		{
			name:     "negative timestamp",
			input:    "-86400",
			expected: time.Unix(-86400, 0).UTC(),
		},
		{
			name:     "null value",
			input:    "null",
			expected: time.Time{},
		},
		{
			name:        "invalid json",
			input:       "\"not a number\"",
			expectError: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var ut UnixTime
			err := json.Unmarshal([]byte(tt.input), &ut)

			if tt.expectError {
				assert.Error(t, err)
				return
			}

			require.NoError(t, err)
			assert.Equal(t, tt.expected, ut.Time)
			assert.Equal(t, time.UTC, ut.Time.Location())
		})
	}
}

func TestUnixTime_MarshalJSON(t *testing.T) {
	tests := []struct {
		name     string
		input    UnixTime
		expected string
	}{
		{
			name:     "positive timestamp",
			input:    UnixTime{Time: time.Unix(1609459200, 0).UTC()},
			expected: "1609459200",
		},
		{
			name:     "zero timestamp",
			input:    UnixTime{Time: time.Unix(0, 0).UTC()},
			expected: "0",
		},
		{
			name:     "negative timestamp",
			input:    UnixTime{Time: time.Unix(-86400, 0).UTC()},
			expected: "-86400",
		},
		{
			name:     "zero time marshals as null",
			input:    UnixTime{Time: time.Time{}},
			expected: "null",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, err := json.Marshal(tt.input)
			require.NoError(t, err)
			assert.Equal(t, tt.expected, string(result))
		})
	}
}

func TestUnixTime_InStruct(t *testing.T) {
	type TestStruct struct {
		ID        int      `json:"id"`
		CreatedAt UnixTime `json:"created_at"`
	}

	t.Run("unmarshal with timestamps", func(t *testing.T) {
		jsonData := []byte(`{"id": 123, "created_at": 1609459200}`)
		var ts TestStruct
		err := json.Unmarshal(jsonData, &ts)
		require.NoError(t, err)

		assert.Equal(t, 123, ts.ID)
		assert.Equal(t, time.Unix(1609459200, 0).UTC(), ts.CreatedAt.Time)
	})

	t.Run("marshal with timestamps", func(t *testing.T) {
		ts := TestStruct{
			ID:        456,
			CreatedAt: UnixTime{Time: time.Unix(1609459200, 0).UTC()},
		}

		jsonData, err := json.Marshal(ts)
		require.NoError(t, err)

		var result map[string]interface{}
		err = json.Unmarshal(jsonData, &result)
		require.NoError(t, err)

		assert.Equal(t, float64(456), result["id"])
		assert.Equal(t, float64(1609459200), result["created_at"])
	})
}

func TestUnixTime_TimezoneBehavior(t *testing.T) {
	loc, err := time.LoadLocation("America/New_York")
	require.NoError(t, err)

	localTime := time.Date(2021, 1, 1, 12, 0, 0, 0, loc)
	ut := UnixTime{Time: localTime}

	jsonData, err := json.Marshal(ut)
	require.NoError(t, err)

	var unmarshaled UnixTime
	err = json.Unmarshal(jsonData, &unmarshaled)
	require.NoError(t, err)

	assert.Equal(t, time.UTC, unmarshaled.Time.Location())
	assert.Equal(t, localTime.Unix(), unmarshaled.Unix())
}
