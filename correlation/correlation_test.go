package correlation

import "testing"
import "github.com/kamgru/gonums/numerics"

func TestCorrelation(t *testing.T) {

	xSamples := []float64{20.0, 30.0, 40.0, 50.0}
	ySamples := []float64{1500.0, 3000.0, 5000.0, 7500.0}

	expected := 0.99380799

	actual := Pearson(xSamples, ySamples)

	if numerics.Round(actual, 8) != expected {
		t.Fail()
	}

}