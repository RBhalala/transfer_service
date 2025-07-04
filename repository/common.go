package repository

import (
	"fmt"
	"math/big"
)

// ParsePositiveDecimal converts a string to *big.Float and enforces > 0.
func ParsePositiveDecimal(s string) (*big.Float, error) {
	f, _, err := big.ParseFloat(s, 10, 256, big.ToNearestEven)
	if err != nil {
		return nil, fmt.Errorf("invalid decimal: %w", err)
	}
	if f.Cmp(big.NewFloat(0)) <= 0 {
		return nil, fmt.Errorf("amount must be positive")
	}
	return f, nil
}

// ParseNonNegativeDecimal converts a string to *big.Float and enforces ≥ 0.
func ParseNonNegativeDecimal(s string) (*big.Float, error) {
	f, _, err := big.ParseFloat(s, 10, 256, big.ToNearestEven)
	if err != nil {
		return nil, fmt.Errorf("invalid decimal: %w", err)
	}
	if f.Cmp(big.NewFloat(0)) < 0 {
		return nil, fmt.Errorf("amount must be positive")
	}
	return f, nil
}

// ParseDecimal is internal‑use: no validation.
func ParseDecimal(s string) (*big.Float, error) {
	f, _, err := big.ParseFloat(s, 10, 256, big.ToNearestEven)
	if err != nil {
		return nil, err
	}
	return f, nil
}
