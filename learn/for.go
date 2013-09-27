package main

import "fmt"

func main() {
	for i := 0; i < 10; i++ {
		fmt.Println(i)
	}

	// for 就是 while= =
	index := 0
	for index < 10 {
		fmt.Println(index)
		index = index + 1
	}

	// 死循环= =
	// for {}
}
