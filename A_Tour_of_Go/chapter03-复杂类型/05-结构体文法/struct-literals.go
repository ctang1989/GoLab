package main

import (
	"fmt"
)

type Vertex struct {
	X, Y int
}

var (
	v1 = Vertex{1, 2}  // 类型为Vertex
	v2 = Vertex{X: 1}  // Y:0 被省略
	v3 = Vertex{}      // X:0 和 Y:0
	p  = &Vertex{1, 2} // 类型为 *Vertex
)

func main() {
	fmt.Println(v1, p, v2, v3)
}

/* 运行结果
{1 2} &{1 2} {1 0} {0 0}
*/

// https://tour.go-zh.org/moretypes/5
