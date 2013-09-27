package main

import "fmt"

func add(a int, b int) int {
	return a + b
}

func sub(a, b int) int {
	return a - b
}

func print(str string) {
	fmt.Println(str)
}

// 返回多值
func swap(x, y string) (string, string) {
	return y, x
}

// 命名返回
func getXandY() (X, Y string) {
	X = "I am X"
	Y = "I am Y"
	return
}

func adder() func(int) int {
	sum := 0
	return func(x int) int {
		sum += x
		return sum
	}
}

// 变参, args 为 slice
func printArgs(args ...string) {
	for _, a := range args {
		fmt.Printf("%s ", a)
	}
	fmt.Print("\n")
}

// 函数作为类型
type operator func(x, y int) int

func Executor(x, y int, f operator) int {
	return f(x, y)
}

func main() {
	fmt.Println(add(1, 5)) // 6
	fmt.Println(sub(5, 2)) // 3

	fmt.Println(Executor(2, 2, add)) // 4
	fmt.Println(Executor(1, 2, sub)) // -1

	print("this is a string")
	x, y := "x", "y"
	fmt.Println("x =", x, ", y =", y)
	x, y = swap(x, y)
	fmt.Println("x =", x, ", y =", y)
	X, Y := getXandY()
	fmt.Println(X)
	fmt.Println(Y)

	ThisIsAFunc := func(x, y int) int {
		return (x + y)
	}

	fmt.Println(ThisIsAFunc(1, 5))

	pos, neg := adder(), adder()
	for i := 0; i < 5; i++ {
		fmt.Println(pos(1), neg(-1))
		fmt.Println(pos(1), neg(-1))
	}
	// 1 -1
	// 2 -2
	// ......

	printArgs("first", "second", "third")

}
