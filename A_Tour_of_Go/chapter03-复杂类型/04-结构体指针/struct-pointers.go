package main

import (
	"fmt"
)

type Vertex struct {
	X int
	Y int
}

func main() {
	v := Vertex{1, 2}
	p := &v
	p.X = 1e9
	fmt.Println(v)
}

/* 运行结果
{1000000000 2}
*/

// https://tour.go-zh.org/moretypes/4
