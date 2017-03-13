package main

import (
	"fmt"
	"math"
)

func Sqrt(x float64) float64 {
	z, d := float64(1), float64(1)

	for d > 1e-15 {
		z0 := z
		z = z - (z*z-x)/(2*z)
		d = math.Abs(z - z0)
	}
	return z
}

func main() {
	fmt.Println(Sqrt(10))
	fmt.Println(math.Sqrt(10))
}

// https://nylira.com/a-tour-of-go-solutions/
