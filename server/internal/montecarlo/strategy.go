package montecarlo

import "kelly-game-theory-simulator/server/internal/kelly"

type StrategyFunc func(
	bankroll float64,
	p float64,
	decimalOdds float64,
) float64

func FlatStrategy(
	bankroll float64,
	p float64,
	decimalOdds float64,
) float64 {
	return 0.05
}

func FullKellyStrategy(
	bankroll float64,
	p float64,
	decimalOdds float64,
) float64 {
	fraction, err := kelly.CalculateKelly(p, decimalOdds)

	if err != nil {
		return 0
	}

	if fraction < 0 {
		return 0
	}

	return fraction
}
