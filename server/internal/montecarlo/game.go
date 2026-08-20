package montecarlo

type GameResult struct {
	Rounds        []RoundResult
	FinalBankRoll float64
	Ruined        bool
	RuinedAtRound int
}

func SimulateGame(
	startingBankRoll float64,
	numRounds int,
	p float64,
	decimalOdds float64,
	strategy StrategyFunc,
) GameResult {
	bankroll := startingBankRoll
	rounds := make([]RoundResult, 0, numRounds)

	ruined := false
	ruinedAtRound := -1

	for round := 1; round <= numRounds; round++ {

		stakeFraction := strategy(bankroll, p, decimalOdds)

		result := SimulateRound(
			bankroll,
			stakeFraction,
			decimalOdds,
			p,
		)

		rounds = append(rounds, result)

		bankroll = result.Bankroll

		if result.IsRuined {
			ruined = true
			ruinedAtRound = round
			break
		}
	}

	return GameResult{
		Rounds:        rounds,
		FinalBankRoll: bankroll,
		Ruined:        ruined,
		RuinedAtRound: ruinedAtRound,
	}
}
