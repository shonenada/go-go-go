package main

import "fmt"

// make用于内建类型（map、slice 和channel）的内存分配。new用于各种类型的内存分配。

type Vertex struct {
	X, Y int
}

func main() {
	v := new(Vertex)
	// 使用 new 返回的是指针，并且所有变量初始值为 0
	fmt.Println(v) // &{0, 0}
	v.X, v.Y = 1, 5
	fmt.Println(v) // &{1, 5}
}
