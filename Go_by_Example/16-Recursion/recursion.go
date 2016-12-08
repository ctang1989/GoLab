package main

import "fmt"

func fact(n int) int {
	if n == 0 {
		return 1
	}
	return n * fact(n-1)
}

func main() {
	fmt.Println(fact(7))
}


/*
$ go run recursion.go 
5040
*/

// https://gobyexample.com/recursion