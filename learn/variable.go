package main

import "fmt"

// Global variable should be defined outside any func.
var global_var = "This is a global variable."

func main() {

	// The first kind of assignment.
	var x1 int
	x1 = 1

	// The 2nd kind of assignment.
	var x2 int = 2

	// The 3rd kind.
	var x3 = 3

	// The 4th kind.
	x4 := 4

	// First mode to assign value into mult variables.
	var y1, y2 int
	y1 = 11
	y2 = 12

	// The 2nd mode
	var y3, y4 int = 13, 14

	// The 3th mode
	var y5, y6 = 15, 16

	// The 4th mode
	y7, y8 := 17, 18

	// Assign to two strings.
	str1, str2 := "This is a String.", "This is another String"

	// '_' is a special variable, whose value will be discarded.
	_, num := 1.0, 2.0

	fmt.Println(global_var)
	fmt.Println(x1, x2, x3, x4)
	fmt.Println(y1, y2, y3, y4, y5, y6, y7, y8)
	fmt.Println(str1, str2)
	fmt.Println(num)
}
