package main

// 分组声明
import (
	"fmt"
)

// 相当于:
// import "fmt"
// import "os"

func main() {
	const (
		pi     = 3.1415926
		prefix = "Go_"
	)
	fmt.Printf("pi = %f\nprefix = %s\n", pi, prefix)

	var (
		i        int = 2
		fileName string
	)
	fmt.Printf("i = %d\nfileName = %s\n", i, fileName)

	// iota: 第一次调用是 0，之后每次调用都增加 1。
	const (
		x = iota // x == 0
		y = iota // y == 1
		z = iota // z == 2
		w        // 在 const 中，如果不对常量赋值，默认是和前一个常量一样的值（z = iota），所以这里 w = 3
	)
	fmt.Printf("x = %d, y = %d, z = %d, w = %d\n", x, y, z, w)

	const v = iota // v == 0
	// iota 在每个 const 中，都会刷新一遍，即从 0 开始递增。
	fmt.Printf("v = %d\n", v)

	const (
		e, f, g = iota, iota, iota
		// e=0, f=0, g=0    iota在同一行值相同
	)
	fmt.Printf("e = %d, f = %d, g = %d\n", e, f, g)
}
