// Copyright © 2025 Ping Identity Corporation

package types_test

import (
	"encoding/json"
	"math/big"
	"testing"

	"github.com/pingidentity/pingone-go-client/types"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestBigFloat_MarshalJSON(t *testing.T) {
	tests := []struct {
		name     string
		input    types.BigFloatUnquoted
		expected string
		wantErr  bool
	}{
		{
			name: "positive integer",
			input: types.BigFloatUnquoted{
				Float: big.NewFloat(123),
			},
			expected: "1.23e+02",
			wantErr:  false,
		},
		{
			name: "negative integer",
			input: types.BigFloatUnquoted{
				Float: big.NewFloat(-456),
			},
			expected: "-4.56e+02",
			wantErr:  false,
		},
		{
			name: "positive decimal",
			input: types.BigFloatUnquoted{
				Float: big.NewFloat(123.456),
			},
			expected: "1.23456e+02",
			wantErr:  false,
		},
		{
			name: "negative decimal",
			input: types.BigFloatUnquoted{
				Float: big.NewFloat(-789.012),
			},
			expected: "-7.89012e+02",
			wantErr:  false,
		},
		{
			name: "zero",
			input: types.BigFloatUnquoted{
				Float: big.NewFloat(0),
			},
			expected: "0e+00",
			wantErr:  false,
		},
		{
			name: "very small number",
			input: types.BigFloatUnquoted{
				Float: big.NewFloat(0.000001),
			},
			expected: "1e-06",
			wantErr:  false,
		},
		{
			name: "very large number",
			input: types.BigFloatUnquoted{
				Float: big.NewFloat(123456789012345),
			},
			expected: "1.23456789012345e+14",
			wantErr:  false,
		},
		{
			name: "nil float",
			input: types.BigFloatUnquoted{
				Float: nil,
			},
			expected: "null",
			wantErr:  false,
		},
		{
			name: "scientific notation input",
			input: func() types.BigFloatUnquoted {
				f, _, _ := big.ParseFloat("1.5e10", 10, 53, big.ToNearestEven)
				return types.BigFloatUnquoted{Float: f}
			}(),
			expected: "1.5e+10",
			wantErr:  false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, err := tt.input.MarshalJSON()

			if tt.wantErr {
				assert.Error(t, err)
				return
			}

			assert.NoError(t, err)
			assert.Equal(t, tt.expected, string(result))
		})
	}
}

func TestBigFloat_UnmarshalJSON(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected func() *big.Float
		wantErr  bool
		errMsg   string
	}{
		{
			name:  "positive integer",
			input: "123",
			expected: func() *big.Float {
				return big.NewFloat(123)
			},
			wantErr: false,
		},
		{
			name:  "negative integer",
			input: "-456",
			expected: func() *big.Float {
				return big.NewFloat(-456)
			},
			wantErr: false,
		},
		{
			name:  "positive decimal",
			input: "123.456",
			expected: func() *big.Float {
				return big.NewFloat(123.456)
			},
			wantErr: false,
		},
		{
			name:  "negative decimal",
			input: "-789.012",
			expected: func() *big.Float {
				return big.NewFloat(-789.012)
			},
			wantErr: false,
		},
		{
			name:  "zero",
			input: "0",
			expected: func() *big.Float {
				return big.NewFloat(0)
			},
			wantErr: false,
		},
		{
			name:  "scientific notation positive exponent",
			input: "1.23e+02",
			expected: func() *big.Float {
				f, _, _ := big.ParseFloat("1.23e+02", 10, 53, big.ToNearestEven)
				return f
			},
			wantErr: false,
		},
		{
			name:  "scientific notation negative exponent",
			input: "1.5e-06",
			expected: func() *big.Float {
				f, _, _ := big.ParseFloat("1.5e-06", 10, 53, big.ToNearestEven)
				return f
			},
			wantErr: false,
		},
		{
			name:  "null value",
			input: "null",
			expected: func() *big.Float {
				return nil
			},
			wantErr: false,
		},
		{
			name:    "invalid format - letters",
			input:   "abc",
			wantErr: true,
			errMsg:  "BigFloat: could not parse",
		},
		{
			name:    "invalid format - mixed",
			input:   "12.34.56",
			wantErr: true,
			errMsg:  "BigFloat: could not parse",
		},
		{
			name:    "invalid format - empty string",
			input:   "",
			wantErr: true,
			errMsg:  "BigFloat: could not parse",
		},
		{
			name:  "very large number",
			input: "123456789012345",
			expected: func() *big.Float {
				f, _, _ := big.ParseFloat("123456789012345", 10, 53, big.ToNearestEven)
				return f
			},
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var result types.BigFloatUnquoted
			err := result.UnmarshalJSON([]byte(tt.input))

			if tt.wantErr {
				assert.Error(t, err)
				if tt.errMsg != "" {
					assert.Contains(t, err.Error(), tt.errMsg)
				}
				return
			}

			assert.NoError(t, err)
			expected := tt.expected()

			if expected == nil {
				assert.Nil(t, result.Float)
			} else {
				require.NotNil(t, result.Float)
				assert.Equal(t, expected.Text('e', -1), result.Text('e', -1))
			}
		})
	}
}

func TestBigFloat_JSONRoundTrip(t *testing.T) {
	tests := []struct {
		name  string
		value *big.Float
	}{
		{
			name:  "positive integer",
			value: big.NewFloat(42),
		},
		{
			name:  "negative integer",
			value: big.NewFloat(-42),
		},
		{
			name:  "positive decimal",
			value: big.NewFloat(3.14159),
		},
		{
			name:  "negative decimal",
			value: big.NewFloat(-2.71828),
		},
		{
			name:  "zero",
			value: big.NewFloat(0),
		},
		{
			name:  "very small number",
			value: big.NewFloat(0.000000001),
		},
		{
			name:  "very large number",
			value: big.NewFloat(999999999999),
		},
		{
			name:  "nil value",
			value: nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			original := types.BigFloatUnquoted{Float: tt.value}

			// Marshal to JSON
			jsonBytes, err := json.Marshal(original)
			require.NoError(t, err)

			// Unmarshal from JSON
			var decoded types.BigFloatUnquoted
			err = json.Unmarshal(jsonBytes, &decoded)
			require.NoError(t, err)

			// Compare values
			if tt.value == nil {
				assert.Nil(t, decoded.Float)
			} else {
				require.NotNil(t, decoded.Float)
				// Compare string representations in scientific notation
				assert.Equal(t, original.Text('e', -1), decoded.Text('e', -1))
			}
		})
	}
}

func TestBigFloat_InStruct(t *testing.T) {
	type TestStruct struct {
		ID    string                  `json:"id"`
		Value types.BigFloatUnquoted  `json:"value"`
		Score *types.BigFloatUnquoted `json:"score,omitempty"`
	}

	tests := []struct {
		name     string
		input    string
		expected TestStruct
		wantErr  bool
	}{
		{
			name:  "complete struct with values",
			input: `{"id":"test-1","value":1.23e+02,"score":4.56e+01}`,
			expected: TestStruct{
				ID:    "test-1",
				Value: types.BigFloatUnquoted{Float: big.NewFloat(123)},
				Score: &types.BigFloatUnquoted{Float: big.NewFloat(45.6)},
			},
			wantErr: false,
		},
		{
			name:  "struct with null value",
			input: `{"id":"test-2","value":null}`,
			expected: TestStruct{
				ID:    "test-2",
				Value: types.BigFloatUnquoted{Float: nil},
			},
			wantErr: false,
		},
		{
			name:  "struct with omitted optional field",
			input: `{"id":"test-3","value":7.89e+00}`,
			expected: TestStruct{
				ID:    "test-3",
				Value: types.BigFloatUnquoted{Float: big.NewFloat(7.89)},
			},
			wantErr: false,
		},
		{
			name:  "struct with zero value",
			input: `{"id":"test-4","value":0e+00}`,
			expected: TestStruct{
				ID:    "test-4",
				Value: types.BigFloatUnquoted{Float: big.NewFloat(0)},
			},
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var result TestStruct
			err := json.Unmarshal([]byte(tt.input), &result)

			if tt.wantErr {
				assert.Error(t, err)
				return
			}

			require.NoError(t, err)
			assert.Equal(t, tt.expected.ID, result.ID)

			// Compare BigFloat values
			if tt.expected.Value.Float == nil {
				assert.Nil(t, result.Value.Float)
			} else {
				require.NotNil(t, result.Value.Float)
				assert.Equal(t, tt.expected.Value.Text('e', -1), result.Value.Float.Text('e', -1))
			}

			// Compare optional Score field
			if tt.expected.Score == nil {
				assert.Nil(t, result.Score)
			} else {
				require.NotNil(t, result.Score)
				if tt.expected.Score.Float == nil {
					assert.Nil(t, result.Score.Float)
				} else {
					require.NotNil(t, result.Score.Float)
					assert.Equal(t, tt.expected.Score.Text('e', -1), result.Score.Float.Text('e', -1))
				}
			}
		})
	}
}

func TestBigFloat_StructRoundTrip(t *testing.T) {
	type TestStruct struct {
		ID    string                 `json:"id"`
		Value types.BigFloatUnquoted `json:"value"`
	}

	original := TestStruct{
		ID:    "round-trip-test",
		Value: types.BigFloatUnquoted{Float: big.NewFloat(123.456)},
	}

	// Marshal to JSON
	jsonBytes, err := json.Marshal(original)
	require.NoError(t, err)

	// Unmarshal from JSON
	var decoded TestStruct
	err = json.Unmarshal(jsonBytes, &decoded)
	require.NoError(t, err)

	// Compare values
	assert.Equal(t, original.ID, decoded.ID)
	require.NotNil(t, decoded.Value.Float)
	assert.Equal(t, original.Value.Text('e', -1), decoded.Value.Text('e', -1))
}
