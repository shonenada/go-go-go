package main

import "fmt"

type complexNum struct {
	real  int
	image int
}

func main() {
	c := complexNum{1, 2}
	p := &c
	p.real = 5
	fmt.Println(c)
}
