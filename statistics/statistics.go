package statistics

import (
    "github.com/kamgru/gonums/numerics"
    "math"
)

func PopulationVariance(input []float64) float64 {
    mean := numerics.Average(input)

    parts := make([]float64, len(input))

    for i, n := range input {
        parts[i] = math.Pow(n - mean, 2)
    }

    variance := numerics.Average(parts)

    return variance
}

func SampleVariance(input []float64) float64 {
    mean := numerics.Average(input)

    parts := make([]float64, len(input))

    for i, n := range input {
        parts[i] = math.Pow(n - mean, 2)
    }

    variance := numerics.Sum(parts) / (float64(len(input) - 1))

    return variance
}

func PopulationStandardDeviation(input []float64) float64 {
    standardDeviation := math.Sqrt(PopulationVariance(input))
    return standardDeviation
}

func SampleStandardDeviation(input []float64) float64 {
    standardDeviation := math.Sqrt(SampleVariance(input))
    return standardDeviation
}