package gonums

import (
	"math"
)

func Avg(input []float64) float64 {
	sum := 0.0
	for _, i := range input {
		sum += i
	}

	avg := sum / float64(len(input))

	return avg
}

func Corr(xSamples, ySamples []float64) float64 {

	sumProducts := 0.0
	sumX := 0.0
	sumY := 0.0
	sumXSq := 0.0
	sumYSq := 0.0

	for i := 0; i < len(xSamples); i++ {
		sumProducts += (xSamples[i] * ySamples[i])
		sumX += xSamples[i]
		sumY += ySamples[i]
		sumXSq += xSamples[i] * xSamples[i]
		sumYSq += ySamples[i] * ySamples[i]
	}

	len := float64(len(xSamples))
	corr := (len*sumProducts - (sumX * sumY)) / (math.Sqrt(len*sumXSq-sumX*sumX) * math.Sqrt(len*sumYSq-sumY*sumY))

	return corr
}

func Round(num float64, decimalCount int) float64 {
	mod := math.Pow10(decimalCount)
	return math.Round(num*mod) / mod
}
