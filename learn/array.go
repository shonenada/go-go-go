package main

import "fmt"

func main() {
	// 声明数组： var variable [size]type
	var arr [5]int
	arr[0] = 1
	arr[1] = 2

	fmt.Println("arr[0] =", arr[0]) // arr[0] = 1
	fmt.Println("arr[1] =", arr[1]) // arr[1] = 2
	fmt.Println("arr[2] =", arr[2]) // arr[2] = 0
	// 数组作为参数传入函数的时候，传入的是副本而不是指针。

	// 使用 := 声明
	array := [5]int{1, 2, 3}
	fmt.Printf("%d, %d, %d, %d, %d\n", array[0], array[1], array[2], array[3], array[4])
	//1, 2, 3, 0, 0

	// 使用 [...] 自动计算长度
	a := [...]int{0, 1, 2}
	fmt.Printf("%d, %d, %d\n", a[0], a[1], a[2])
	// 0, 1, 2

	// 二维数组
	doubleArray := [2][4]int{{0, 1, 2, 3}, {4, 5, 6, 7}}
	i, j := 0, 0
	for i < 2 {
		for j < 4 {
			fmt.Printf("doubleArray[%d][%d] = %d\n", i, j, doubleArray[i][j])
			j++
		}
		i++
	}

	for i, v := range array {
		fmt.Println("index:", i, "value:", v)
	}
}
