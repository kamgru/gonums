package numerics

import "math"

func Average(input []float64) float64 {
	sum := 0.0
	for _, i := range input {
		sum += i
	}

	average := sum / float64(len(input))

	return average
}

func Round(number float64, decimalCount int) float64 {
	decimalModifier := math.Pow10(decimalCount)
	return math.Round(number * decimalModifier) / decimalModifier
}