package main

import (
	"fmt"
)

// error type
type ErrNegativeSqrt float64

// error function for the error type
func (e ErrNegativeSqrt) Error() string {
	return fmt.Sprintf("Cannot Sqrt negative number: %v", float64(e))
}

func Sqrt(x float64) (float64, error) {

	// return error if number is negative
	if x < 0 {
		return 0, ErrNegativeSqrt(x)

		// otherwise find an approximation for the square root
	} else {
		z := float64(1)
		for i := 0; i < 1000; i++ {
			z = z - (z*z-x)/(2*z)
		}
		return z, nil
	}

}

func main() {
	fmt.Println(Sqrt(2))
	fmt.Println(Sqrt(-2))
}

/*
	1.414213562373095 <nil>
	0 Cannot Sqrt negative number: -2
*/
