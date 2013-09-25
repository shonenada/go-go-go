package main

import "fmt"

// 全局变量需要在所有函数外部声明定义
var global_var = "This is a global variable."

func main() {

	// 第一种赋值的方式，可以显示声明变量的类型
	var x1 int
	x1 = 1

	// 第二种赋值的方式，可以直接初始化
	var x2 int = 2

	// 第三种赋值的方式，可以不显示声明类型，编译器自动推导出变量类型
	var x3 = 3

	// 第四种赋值方式，使用 := 操作符
	x4 := 4

	// 赋值给多个变量
	var y1, y2 int
	y1 = 11
	y2 = 12

	// 第二种赋值给多个变量的形式
	var y3, y4 int = 13, 14

	// 第三种
	var y5, y6 = 15, 16

	// 第四种
	y7, y8 := 17, 18

	// 赋值给两个字符串
	str1, str2 := "This is a String.", "This is another String"

	// '_' 是一个特殊的变量，它会抛弃被赋予的值
	_, num := 1.0, 2.0

	fmt.Println(global_var)
	fmt.Println(x1, x2, x3, x4)
	fmt.Println(y1, y2, y3, y4, y5, y6, y7, y8)
	fmt.Println(str1, str2)
	fmt.Println(num)
}
