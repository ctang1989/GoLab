package main

import (
	"fmt"
)

func split(sum int) (x, y int) {
	x = sum * 4 / 9
	y = sum - x
	return
}

func main() {
	fmt.Println(split(17))
}

/* 运行结果
7 10
*/

// https://yushuangqi.com/book/gotour/basics/named-returns.html
