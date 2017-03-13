package main

import (
	"fmt"
)

func swap(x, y string) (string, string) {
	return y, x
}

func main() {
	a, b := swap("hello", "world")
	fmt.Println(a, b)
}

// https://yushuangqi.com/book/gotour/basics/multi_value_return.html
