package gonums

import (
	"testing"
)

func TestRound(t *testing.T) {
	input := 0.98723
	expected := 0.99

	actual := Round(input, 2)

	if actual != expected {
		t.Fail()
	}
}

func TestAvg(t *testing.T) {
	nums := []float64{2.0, 3.0, 4.0}

	res := Avg(nums)

	if res != 3.0 {
		t.Fatalf("expected %v actual %v", 3.0, res)
	}
}

func TestCorrelation(t *testing.T) {

	xSamples := []float64{20.0, 30.0, 40.0, 50.0}
	ySamples := []float64{1500.0, 3000.0, 5000.0, 7500.0}

	expected := 0.99380799

	actual := Corr(xSamples, ySamples)

	if Round(actual, 8) != expected {
		t.Fail()
	}

}
