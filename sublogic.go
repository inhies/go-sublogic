// Package sublogic implements subjective logic functions
package sublogic

import (
	"fmt"
)

type Opinion struct {
	Belief      float64
	Disbelief   float64
	Uncertainty float64
	Baserate    float64
	Expectation float64
}

// Creates a new Opinion strucure after verifying that valid data was supplied
func NewOpinion(belief, disbelief, uncertainty, baserate float64) (opinion *Opinion, err error) {

	// belief + disbelief + uncertainty should always sum to 1
	if belief+disbelief+uncertainty != 1 {
		err = fmt.Errorf("Sum does not equal 1")
		return
	}

	// belief,disbelief,uncertainty, and baserate should all be within the range {0,1}
	if belief < 0 || belief > 1 {
		err = fmt.Errorf("Belief out of range")
		return
	}

	if disbelief < 0 || disbelief > 1 {
		err = fmt.Errorf("Disbelief out of range")
		return
	}

	if uncertainty < 0 || uncertainty > 1 {
		err = fmt.Errorf("Uncertainty out of range")
		return
	}

	if baserate < 0 || baserate > 1 {
		err = fmt.Errorf("Baserate out of range")
		return
	}

	expectation := belief + baserate*uncertainty
	opinion = &Opinion{belief, disbelief, uncertainty, baserate, expectation}
	return
}

// Where A trusts B, and B trusts C, this returns A's transitive trust in C
func (A *Opinion) Discount(B *Opinion) (C *Opinion, err error) {
	C = &Opinion{}

	C.Belief = A.Belief * B.Belief
	C.Disbelief = A.Belief * B.Disbelief
	C.Uncertainty = (A.Disbelief + A.Uncertainty + A.Belief*B.Uncertainty)
	C.Baserate = B.Baserate
	C.Expectation = C.Belief + C.Baserate*C.Uncertainty
	return
}

// Where A trusts C, and B trusts C, this returns the consensus of their trust in C
func (A *Opinion) Fuse(B *Opinion) (AC *Opinion, err error) {
	AC = &Opinion{}

	k := A.Uncertainty + B.Uncertainty - A.Uncertainty*B.Uncertainty
	l := A.Uncertainty + B.Uncertainty - 2*A.Uncertainty*B.Uncertainty
	// Check for a 0 divisor
	if k != 0 {
		// Normal case: Argument opinions are not both vacuous
		if l != 0 {
			AC.Belief = (A.Belief*B.Uncertainty + B.Belief*A.Uncertainty) / k
			AC.Disbelief = (A.Disbelief*B.Uncertainty + B.Disbelief*A.Uncertainty) / k
			AC.Uncertainty = (A.Uncertainty * B.Uncertainty) / k
			AC.Baserate = (B.Baserate*A.Uncertainty + A.Baserate*B.Uncertainty -
				(A.Baserate+B.Baserate)*A.Uncertainty*B.Uncertainty) / l
		} else {
			// Argument opinions are both vacuous
			if A.Uncertainty == B.Uncertainty {
				AC.Belief = 0
				AC.Disbelief = 0
				AC.Uncertainty = 1
				AC.Baserate = A.Baserate
			} else {
				err = fmt.Errorf("Relative atomicities are NOT equal")
				return
			}
		}
	} else {
		// I made an assumption and used gamma = 1 as this appears to mirror the behavior of
		// the simulator: http://folk.uio.no/josang/sl/Op.html although elsewhere I have 
		// seen the gamma value passed in as an argument to the function. I'm not great with 
		// calculus so if someone can shed some more light or correct this please do!
		gamma := 1.0

		AC.Belief = (gamma*A.Belief + B.Belief) / (gamma + 1)
		AC.Disbelief = (gamma*A.Disbelief + B.Disbelief) / (gamma + 1)
		AC.Uncertainty = 0.0
		AC.Baserate = (gamma*B.Baserate + A.Baserate) / (gamma + 1)
	}
	AC.Expectation = AC.Belief + AC.Baserate*AC.Uncertainty
	return
}
