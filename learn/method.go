package main

import (
	"fmt"
)

// Go 中并没有 classes，但是可以使用 method 对 struct 添加函数。形式上类似于 classes

type complexNum struct {
	real, image int
}

// 指针，传入的是对象本身
func (c *complexNum) PrintReal() {
	fmt.Println(c.real)
	c.real = 0
}

// 引用，传入的是对象的副本
func (c complexNum) PrintImage() {
	fmt.Println(c.image)
	c.image = 0
}

func main() {
	c := &complexNum{2, 5}
	c.PrintReal()
	cn := complexNum{1, 3}
	cn.PrintImage()

	fmt.Println(c.real)
	fmt.Println(cn.image)
}
