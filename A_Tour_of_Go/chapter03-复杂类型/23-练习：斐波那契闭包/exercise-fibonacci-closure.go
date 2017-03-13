package main

import (
	"fmt"
)

// fibonacci函数会返回一个返回int的函数。
func fibonacci() func() int {
	first, second := 0, 1
	return func() int {
		first, second = second, first+second
		return first
	}
}

func main() {
	f := fibonacci()
	for i := 0; i < 10; i++ {
		fmt.Println(f())
	}
}

/*
	1
	1
	2
	3
	5
	8
	13
	21
	34
	55
*/
