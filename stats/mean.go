package stats

import "math"

// Mean gets the average of a slice of numbers
func Mean(input []float64) (float64, error) {

	if len(input) == 0 {
		return math.NaN(), EmptyInputErr
	}

	sum, _ := Sum(input)

	return sum / float64(len(input)), nil
}

// GeometricMean gets the geometric mean for a slice of numbers
func GeometricMean(input []float64) (float64, error) {

	l := len(input)
	if l == 0 {
		return math.NaN(), EmptyInputErr
	}

	// Get the product of all the numbers
	var p float64
	for _, n := range input {
		if p == 0 {
			p = n
		} else {
			p *= n
		}
	}

	// Calculate the geometric mean
	return math.Pow(p, 1/float64(l)), nil
}

// HarmonicMean gets the harmonic mean for a slice of numbers
func HarmonicMean(input []float64) (float64, error) {

	l := len(input)
	if l == 0 {
		return math.NaN(), EmptyInputErr
	}

	// Get the sum of all the numbers reciprocals and return an
	// error for values that cannot be included in harmonic mean
	var p float64
	for _, n := range input {
		if n < 0 {
			return math.NaN(), NegativeErr
		} else if n == 0 {
			return math.NaN(), ZeroErr
		}
		p += (1 / n)
	}

	return float64(l) / p, nil
}
