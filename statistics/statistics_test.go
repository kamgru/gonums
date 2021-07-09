package statistics

import (
    "github.com/kamgru/gonums/numerics"
    "testing"
)

func TestPopulationVariance(t *testing.T) {
    input := []float64 {60, 47, 17, 43, 30}
    expected := 217.04
    actual := PopulationVariance(input)

    if numerics.Round(actual, 2) != expected {
        t.Fail()
    }
}

func TestSampleVariance(t *testing.T) {
    input := []float64 {60, 47, 17, 43, 30}
    expected := 271.30
    actual := SampleVariance(input)

    if numerics.Round(actual, 2) != expected {
        t.Fail()
    }
}

func TestPopulationStandardDeviation(t *testing.T) {
    input := []float64 {60, 47, 17, 43, 30}
    expected := 14.73227749

    actual := PopulationStandardDeviation(input)

    if numerics.Round(actual, 8) != expected {
        t.Fail()
    }
}

func TestSampleStandardDeviation(t *testing.T) {
    input := []float64 {60, 47, 17, 43, 30}
    expected := 16.47118696

    actual := SampleStandardDeviation(input)

    if numerics.Round(actual, 8) != expected {
        t.Fail()
    }
}