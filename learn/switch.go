package main

import (
	"fmt"
	"runtime"
	"time"
)

func main() {
	switch os := runtime.GOOS; os {
	case "darwin":
		fmt.Println("OS X.")
	case "linux":
		fmt.Println("Linux")
	default:
		fmt.Println(os)
	}

	// 没有参数的 switch, 当 case 满足条件，后面的语句将不再执行.
	t := time.Now()
	switch {
	case t.Hour() < 12:
		fmt.Println("Morning!")
	case t.Hour() < 17:
		fmt.Println("Afternoon!")
	default:
		fmt.Println("Evening!")
	}

}
