// Copyright © 2025 Ping Identity Corporation

package types

import (
	"fmt"
	"math/big"
)

// BigFloatUnquoted is a wrapper around big.Float to handle JSON marshalling as a string without quotes, as the P1 API expects.
type BigFloatUnquoted struct {
	*big.Float
}

func (b BigFloatUnquoted) MarshalJSON() ([]byte, error) {
	if b.Float == nil {
		return []byte("null"), nil
	}
	return fmt.Appendf(nil, `%s`, b.Text('e', -1)), nil
}

func (b *BigFloatUnquoted) UnmarshalJSON(p []byte) error {
	if string(p) == "null" {
		b.Float = nil
		return nil
	}

	f, _, err := new(big.Float).Parse(string(p), 10)
	if err != nil {
		return fmt.Errorf("BigFloat: could not parse %q", string(p))
	}
	b.Float = f
	return nil
}
