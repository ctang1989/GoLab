package main

import "fmt"

func intSeq() func() int {
	i := 0
	return func() int {
		i += 1
		return i
	}
}

func main() {
	
	nextInt := intSeq()
	fmt.Println(nextInt())
	fmt.Println(nextInt())
	fmt.Println(nextInt())

	newInts := intSeq()
	fmt.Println(newInts())
}


/*
$ go run closures.go 
1
2
3
1
*/

// https://gobyexample.com/closures

// 关于闭包
// http://www.cnblogs.com/Jifangliang/archive/2008/08/05/1260602.html