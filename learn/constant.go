package main

import "fmt"

func main() {

	// Define a string constant.
	const thisFileName string = "constant.go"

	// Define a float32 constant
	const pi float32 = 3.1415926

	// Another mode to define a string constant.
	const prefix = "Go-"

	fmt.Println(thisFileName, pi, prefix)
}
