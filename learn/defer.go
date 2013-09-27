package main

import "fmt"

// defer 是延迟函数，会倒序执行函数，最后返回函数

func dFunc() {
	fmt.Println("the first line")
	defer fmt.Println("the last line")
	defer fmt.Println("the third line")
	fmt.Println("the second line")
}

func main() {
	for i := 0; i < 5; i++ {
		defer fmt.Println(i)
	}
	// 4 3 2 1 0
	fmt.Println()
	dFunc()
}
