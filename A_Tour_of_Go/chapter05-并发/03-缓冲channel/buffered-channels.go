package main

import (
	"fmt"
)

func main() {
	ch := make(chan int, 2)
	ch <- 1
	ch <- 2
	// ch <- 1	// 修改例子使得缓冲区被填满，然后看看会发生什么。
	fmt.Println(<-ch)
	fmt.Println(<-ch)
}

/*
	fatal error: all goroutines are asleep - deadlock!

	goroutine 1 [chan send]:
	main.main()
		F:/Git/GoLab/A_Tour_of_Go/chapter05-并发/03-缓冲channel/buffered-channels.go:11 +0xdb
*/
