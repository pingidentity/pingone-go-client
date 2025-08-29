// Copyright © 2025 Ping Identity Corporation

package pingone

import (
	"fmt"
	"math/big"
)

// BigFloat is a wrapper around big.Float to handle JSON marshalling as a string without losing precision.
type BigFloat struct {
	*big.Float
}

func (b BigFloat) MarshalJSON() ([]byte, error) {
	if b.Float == nil {
		return []byte("null"), nil
	}
	return []byte(fmt.Sprintf(`%s`, b.Text('e', -1))), nil
}

func (b *BigFloat) UnmarshalJSON(p []byte) error {
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
