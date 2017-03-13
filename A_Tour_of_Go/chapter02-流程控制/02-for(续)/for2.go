package main

import (
	"fmt"
)

func main() {
	sum := 1
	for ; sum < 1000; {
		sum += sum
	}
	fmt.Println(sum)
}

// 循环初始化语句和后置语句都是可选的。
// https://yushuangqi.com/book/gotour/flowcontrol/for-continued.html
