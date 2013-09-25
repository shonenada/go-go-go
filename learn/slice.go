package main

import "fmt"

// slice 跟 Python 的 slice 类似= =

func main() {
	// 声明 slice，[] 内不需要任何内容
	var fslice []int // fslice == nil
	slice := []byte{'a', 'b', 'c', 'd', 'e', 'f'}
	fmt.Println("fslice = ", fslice)
	fmt.Println("fslice == nul is", fslice == nil)
	fmt.Println("slice = ", slice)

	// 使用 make 函数创建 slice
	mslice := make([]int, 5) // len = 5
	fmt.Println(mslice)

	mcslice := make([]int, 0, 5) // len = 0, cap = 5
	fmt.Println(mcslice)
	mcslice = append(mcslice, 0)
	mcslice = append(mcslice, 1)
	fmt.Println(mcslice)

	// 从 array 中获得 slice
	var aslice []byte
	arr := []byte{'a', 'b', 'c', 'd', 'e', 'f'}
	aslice = arr[2:4]
	fmt.Println("aslice = ", aslice)

	// 从 slice 中获得 slice
	sslice := slice[2:4]
	fmt.Println("sslice =", sslice)

	// arr[:n] = arr[0:n]
	// arr[n:] = arr[n:len(arr)]

	// slice 的内置函数
	fmt.Println("len(sslice) =", len(sslice)) // 长度
	fmt.Println("cap(sslice) =", cap(sslice)) // 最大容量，slice 开始位置到数组最后位置的长度
	nslice := append(sslice, 'x', 'y', 'z')   // append
	fmt.Println("nslice =", nslice)
	var nnslice []byte
	copy(sslice, nnslice)
	fmt.Println("nnslice =", nnslice)
}
