package main

import (
	"fmt"
	"golang.org/x/tour/tree"
)

// Walk步进tree t将所有的值从tree发送到channel ch.
func Walk(t *tree.Tree, ch chan int) {
	walkRecur(t, ch)
	close(ch)
}

func walkRecur(t *tree.Tree, ch chan int) {
	if t == nil {
		return
	}
	walkRecur(t.Left, ch)
	ch <- t.Value
	walkRecur(t.Right, ch)
}

// Same 检测树t1和t2是否含有相同的值。
func Same(t1, t2 *tree.Tree) bool {
	ch1 := make(chan int)
	go Walk(t1, ch1)
	ch2 := make(chan int)
	go Walk(t2, ch2)
	for n := range ch1 {
		if n != <-ch2 {
			return false
		}
	}
	return true
}

func main() {
	ch := make(chan int)
	go Walk(tree.New(2), ch)
	for n := range ch {
		fmt.Printf("%v ", n)
	}
	fmt.Println()
	fmt.Println(Same(tree.New(1), tree.New(1)))
	fmt.Println(Same(tree.New(1), tree.New(2)))
}

/*
2 4 6 8 10 12 14 16 18 20
true
false
*/
