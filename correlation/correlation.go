package correlation

import (
	"math"
)

func Pearson(xSamples, ySamples []float64) float64 {

	sumProducts := 0.0
	sumX := 0.0
	sumY := 0.0
	sumXSq := 0.0
	sumYSq := 0.0

	for i := 0; i < len(xSamples); i++ {
		sumProducts += xSamples[i] * ySamples[i]
		sumX += xSamples[i]
		sumY += ySamples[i]
		sumXSq += xSamples[i] * xSamples[i]
		sumYSq += ySamples[i] * ySamples[i]
	}

	sampleCount := float64(len(xSamples))
	correlationCoefficient := (sampleCount*sumProducts - (sumX * sumY)) /
		(math.Sqrt(sampleCount * sumXSq-sumX * sumX) * math.Sqrt(sampleCount * sumYSq-sumY * sumY))

	return correlationCoefficient
}
