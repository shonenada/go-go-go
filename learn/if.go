package main

import "fmt"

func main() {
	var IsPrint bool
	if IsPrint {
		fmt.Println("IsPrint = true") // 不会输出
	}
	IsPrint = true
	if IsPrint {
		fmt.Println("IsPrint = true") // 输出
	}

	// 附带语句的 if 语句
	if sum := 5; sum < 6 {
		fmt.Println(sum) // 5
	}

	if !IsPrint {
		fmt.Println("IsPrint = false")
	} else {
		fmt.Println("IsPrint = true")
	}

}
