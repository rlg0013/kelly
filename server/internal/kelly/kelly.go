package kelly

import "fmt"

func CalculateKelly(p float64, decimalOdds float64) (float64, error) {
	if p < 0 || p > 1 {
		return 0, fmt.Errorf("probability must be between 0 and 1")
	}

	if decimalOdds <= 1 {
		return 0, fmt.Errorf("decimal odds must be greater than 1")
	}

	b := decimalOdds - 1
	q := 1 - p

	fraction := (b*p - q) / b

	return fraction, nil
}
