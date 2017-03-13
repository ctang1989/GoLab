package main

import (
	"fmt"
)

func main() {
	pow := make([]int, 10)
	for i := range pow {
		pow[i] = 1 << uint(i) // (x << n == x*2^n ) (x >> n == x*2^(-n))
	}
	for _, value := range pow {
		fmt.Printf("%d\n", value)
	}
}
