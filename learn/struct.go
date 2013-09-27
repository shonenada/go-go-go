package main

import "fmt"

type complexNum struct {
	real  int
	image int
}

// Struct 匿名字段
type People struct {
	name   string
	gender string
}

type Student struct {
	People // 匿名字段
	stu_no string
}

func main() {
	fmt.Println(complexNum{1, 2})

	cnum := complexNum{1, 2}
	cnum.real = 5
	fmt.Println(cnum)

	var (
		c = complexNum{2, 5}
		p = &complexNum{3, 5}
		r = complexNum{image: 5}
		s = complexNum{}
	)

	fmt.Println(c) // {2, 5}
	fmt.Println(p) // &{3, 5}
	fmt.Println(r) // {0, 5}
	fmt.Println(s) // {0, 0}

	stu := Student{People{"name", "boy"}, "123456"}
	fmt.Println(stu)
	fmt.Println(stu.name) // 直接调用
	fmt.Println(stu.gender)
	fmt.Println(stu.stu_no)
}
