package montecarlo

import "math/rand/v2"

type RoundResult struct {
	Won      bool
	Stake    float64
	Profit   float64
	Bankroll float64
	IsRuined bool
}

func SimulateRound(
	bankroll float64,
	stakeFraction float64,
	decimalOdds float64,
	winProbability float64,
) RoundResult {
	if stakeFraction < 0 {
		stakeFraction = 0
	}

	if stakeFraction > 1 {
		stakeFraction = 1
	}

	stake := stakeFraction * bankroll

	won := rand.Float64() < winProbability

	var profit float64

	if won {
		profit = stake * (decimalOdds - 1)
	} else {
		profit = -stake
	}

	newBankRoll := bankroll + profit

	if newBankRoll < 0 {
		newBankRoll = 0
	}

	isRuined := newBankRoll == 0

	return RoundResult{
		Won:      won,
		Stake:    stake,
		Profit:   profit,
		Bankroll: newBankRoll,
		IsRuined: isRuined,
	}
}
