package numerics

import "testing"

func TestAverage(t *testing.T) {
	input := []float64{2.0, 3.0, 4.0}

	expected := 3.0
	actual := Average(input)

	if actual != expected {
		t.Fatalf("expected %v actual %v", expected, actual)
	}
}

func TestRound(t *testing.T) {
	input := 0.98723
	expected := 0.99

	actual := Round(input, 2)

	if actual != expected {
		t.Fail()
	}
}

func TestVariance(t *testing.T) {
	input := []float64 {60, 47, 17, 43, 30}
	expected := 217.04
	actual := Variance(input)

	if Round(actual, 2) != expected {
		t.Fail()
	}
}

func TestStandardDeviation(t *testing.T) {
	input := []float64 {60, 47, 17, 43, 30}
	expected := 14.73227749

	actual := StandardDeviation(input)

	if Round(actual, 8) != expected {
		t.Fail()
	}
}