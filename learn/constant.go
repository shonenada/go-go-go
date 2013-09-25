package main

import "fmt"

func main() {

	// 定义一个 string 常量
	const thisFileName string = "constant.go"

	// 定义一个 float32 常量
	const pi float32 = 3.1415926

	// 另一种定义常量的模式
	const prefix = "Go-"

	fmt.Println(thisFileName, pi, prefix)
}
