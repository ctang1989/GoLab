package main

import (
	"fmt"
	"math"
)

type Abser interface {
	Abs() float64
}

func main() {
	var a Abser
	f := MyFloat(-math.Sqrt2)
	v := Vertex{3, 4}

	a = f  // a MyFloat 实现了 Abser
	a = &v // a *Vertex 实现了Abser

	// 下面一行，v是一个Vertex（而不是 *Vertex）
	// 所以没有实现 Abser
	a = v

	fmt.Println(a.Abs())
}

type MyFloat float64

func (f MyFloat) Abs() float64 {
	if f < 0 {
		return float64(-f)
	}
	return float64(f)
}

type Vertex struct {
	X, Y float64
}

func (v *Vertex) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

// 接口类型是由一组方法定义的集合。

/*调试步骤
1. 直接运行
.\interfaces.go:22: cannot use v (type Vertex) as type Abser in assignment:
	Vertex does not implement Abser (Abs method has pointer receiver)

2. 注释掉 line 22
	5

3. 注释掉 line 18 和 line 15
	1.4142135623730951
*/

// https://tour.go-zh.org/methods/4
