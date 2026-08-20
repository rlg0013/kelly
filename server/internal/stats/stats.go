package stats

import (
	"kelly-game-theory-simulator/server/internal/montecarlo"
	"sort"
)

func MeanFinalBankroll(results []montecarlo.GameResult) float64 {
	meanBankroll := 0.0
	for _, result := range results {
		meanBankroll += result.FinalBankRoll
	}
	meanBankroll /= float64(len(results))
	return meanBankroll
}

func MedianFinalBankroll(results []montecarlo.GameResult) float64 {
	sorted := make([]float64, 0, len(results))
	for _, result := range results {
		sorted = append(sorted, result.FinalBankRoll)
	}
	sort.Float64s(sorted)
	length := len(sorted)

	var medianBankroll float64

	if length%2 != 0 {
		index := length / 2
		medianBankroll = sorted[index]
	} else {
		index1 := (length / 2) - 1
		index2 := (length / 2)
		medianBankroll = (sorted[index1] + sorted[index2]) / 2
	}

	return medianBankroll
}
