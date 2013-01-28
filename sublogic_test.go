package sublogic

import (
	"testing"
)

// TODO(inhies): Convert these in to table based tests. I know it's ugly :'(

// Checks to make sure the opinion fields all add up to one when creating it
func Test_Opinion(t *testing.T) {
	opinion, err := NewOpinion(1, 0.1, 0.2, 0.5) // Sum is more than 1

	if err == nil {
		t.Error("Opinion creation should have failed!", opinion.Belief, opinion.Disbelief,
			opinion.Uncertainty, opinion.Baserate)
	}
}

// Tests the discount operator
func Test_Discount(t *testing.T) {
	opinionA, _ := NewOpinion(0.69, 0.07, 0.24, 0.62) // A's opinion of B            0.84
	opinionB, _ := NewOpinion(0.51, 0.21, 0.28, 0.54) // B's opinion of C            0.66
	opinionC, _ := opinionA.Discount(opinionB)        // A's derived opinion of C    0.62
	if err := opinionC.CheckConsistency(); err != nil {
		t.Error(err)
	}
	// Round everything to two decimal places
	opinionC.Round(2)

	if (opinionC.Belief != 0.35) || (opinionC.Disbelief != 0.14) ||
		(opinionC.Uncertainty != 0.50) || (opinionC.Baserate != 0.54) ||
		(opinionC.Expectation != 0.62) {
		t.Error("Opinion discount calculation was not correct!", opinionC.Belief,
			opinionC.Disbelief, opinionC.Uncertainty, opinionC.Baserate,
			opinionC.Expectation)
	}
}

// Test the normal case of fuse
func Test_Fuse(t *testing.T) {
	opinionA, _ := NewOpinion(0.49, 0.24, 0.27, 0.41) // A's opinion of X            0.60
	opinionB, _ := NewOpinion(0.30, 0.10, 0.60, 0.47) // B's opinion of X            0.58
	opinionC, _ := opinionA.Fuse(opinionB)            // Derived opinion of X        0.63
	if err := opinionC.CheckConsistency(); err != nil {
		t.Error(err)
	}
	// Round everything to two decimal places
	opinionC.Round(2)

	if (opinionC.Belief != 0.53) || (opinionC.Disbelief != 0.24) ||
		(opinionC.Uncertainty != 0.23) || (opinionC.Baserate != 0.42) ||
		(opinionC.Expectation != 0.63) {
		t.Error("Opinion fuse calculation was not correct!", opinionC.Belief,
			opinionC.Disbelief, opinionC.Uncertainty, opinionC.Baserate,
			opinionC.Expectation)
	}
}

// Test fuse with zero denominator
func Test_Fuse0(t *testing.T) {
	opinionA, _ := NewOpinion(0.63, 0.37, 0.00, 0.34) // A's opinion of X            0.63
	opinionB, _ := NewOpinion(0.66, 0.34, 0.00, 0.48) // B's opinion of X            0.66
	opinionC, _ := opinionA.Fuse(opinionB)            // Derived opinion of X        0.65
	if err := opinionC.CheckConsistency(); err != nil {
		t.Error(err)
	}
	// Round everything to two decimal places
	opinionC.Round(2)

	if (opinionC.Belief != 0.65) || (opinionC.Disbelief != 0.36) ||
		(opinionC.Uncertainty != 0.00) || (opinionC.Baserate != 0.41) ||
		(opinionC.Expectation != 0.65) {
		t.Error("Opinion fuse calculation was not correct!", opinionC.Belief,
			opinionC.Disbelief, opinionC.Uncertainty, opinionC.Baserate,
			opinionC.Expectation)
	}
}

// Test with 1 for uncertanties
func Test_Fuse1(t *testing.T) {
	opinionA, _ := NewOpinion(0.00, 0.00, 1, 0.42) // A's opinion of X            0.42
	opinionB, _ := NewOpinion(0.00, 0.00, 1, 0.18) // B's opinion of X            0.18
	opinionC, _ := opinionA.Fuse(opinionB)         // Derived opinion of X        0.00
	if err := opinionC.CheckConsistency(); err != nil {
		t.Error(err)
	}
	// Round everything to two decimal places
	opinionC.Round(2)

	if (opinionC.Belief != 0.00) || (opinionC.Disbelief != 0.00) ||
		(opinionC.Uncertainty != 1.00) || (opinionC.Baserate != 0.42) ||
		(opinionC.Expectation != 0.42) {
		t.Error("Opinion fuse calculation was not correct!", opinionC.Belief,
			opinionC.Disbelief, opinionC.Uncertainty, opinionC.Baserate,
			opinionC.Expectation)
	}
}
