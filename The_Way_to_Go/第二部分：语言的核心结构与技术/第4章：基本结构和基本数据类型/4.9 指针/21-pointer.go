package main

import (
	"fmt"
)

func main() {
	var i1 = 5
	fmt.Printf("An integer: %d, its location in memory: %p\n", i1, &i1)
	var intP *int
	intP = &i1
	fmt.Printf("The value at memory location %p is %d\n", intP, *intP)
}

/*
&: 取地址符，放到一个变量前使用就会返回相应变量的内存地址。
*: 在指针类型前面加上 * 号（前缀）来获取指针所指向的内容。
*/
