package sublogic

import (
	"fmt"
	"math"
)

const tolerance = 1.0e-10

// Rounds all fields in the opinion to prec digits after the decimal.
func (this *Opinion) Round(prec int) {
	this.Belief = round(this.Belief, prec)
	this.Disbelief = round(this.Disbelief, prec)
	this.Uncertainty = round(this.Uncertainty, prec)
	this.Baserate = round(this.Baserate, prec)
	this.Expectation = round(this.Expectation, prec)
}

// Rounds a single float64 to prec digits after the decimal.
func round(x float64, prec int) float64 {
	if math.IsNaN(x) || math.IsInf(x, 0) {
		return x
	}

	sign := 1.0
	if x < 0 {
		sign = -1
		x *= -1
	}

	var rounder float64
	pow := math.Pow(10, float64(prec))
	intermed := x * pow
	_, frac := math.Modf(intermed)

	if frac >= 0.5 {
		rounder = math.Ceil(intermed)
	} else {
		rounder = math.Floor(intermed)
	}

	return rounder / pow * sign
}

// CheckConsistency ensures that upon adding belief, disbelief, and uncertainty together
// that the total is within our tolerance of 1.
func (this *Opinion) CheckConsistency() error {
	if math.Abs(this.Belief+this.Disbelief+this.Uncertainty-1) > tolerance {
		return fmt.Errorf("Sum does not equal 1")
	}
	return nil
}
