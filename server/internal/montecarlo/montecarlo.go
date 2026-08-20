package montecarlo

func RunMonteCarlo(
	numSimulations int,
	startingBankroll float64,
	numRounds int,
	p float64,
	decimalOdds float64,
	strategy StrategyFunc,
) []GameResult {
	simulations := make([]GameResult, 0, numSimulations)

	for simulation := 1; simulation <= numSimulations; simulation++ {
		result := SimulateGame(startingBankroll, numRounds, p, decimalOdds, strategy)
		simulations = append(simulations, result)

	}

	return simulations
}
