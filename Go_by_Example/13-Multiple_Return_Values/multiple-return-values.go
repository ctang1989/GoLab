package main

import "fmt"

func vals() (int, int) {
	return 3, 7
}

func main() {

	a, b := vals()
	fmt.Println(a)
	fmt.Println(b)

	_, c := vals()
	fmt.Println(c)
}


/*
$ go run multiple-return-values.go 
3
7
7
*/

// https://gobyexample.com/multiple-return-values