package main

import (
	"fmt"
	"math"
	"strconv"
)

// Go 通过 interface 实现了 duck-typing

type Abser interface {
	Abs() float64
}

type MyFloat float64

type Vertex struct {
	X, Y float64
}

func (f MyFloat) Abs() float64 {
	if f < 0 {
		return float64(-f)
	} else {
		return float64(f)
	}
}

func (v *Vertex) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

// fmt 中有这样的定义：
// type Stringer interface {
//      String() string
// }
// 任何实现了 String 方法的类型都能被 fmt.Print 调用
func (v Vertex) String() string {
	return strconv.FormatFloat(v.X, 'f', 2, 64) + " + " + strconv.FormatFloat(v.Y, 'f', 2, 64) + "j "
}

type Element interface{}
type List []Element

func main() {
	var abser Abser
	f := MyFloat(-math.Sqrt2)
	v := Vertex{3, 4}
	// Vertex 实现了 String() 方法，
	fmt.Println("v =", v)

	abser = f // a Float32 implements Abser
	fmt.Println("abser = f: ", abser.Abs())

	abser = &v // a *Vertex implements Abser
	fmt.Println("abser = &v:", abser.Abs())
	// abser = v  // a Vertex, does NOT; (Vertex does not implement Abser (Abs method has pointer receiver)

	// 空接口可以用来存放任何数据
	var store interface{}
	int_v := 200
	string_v := "string stored in interface{}"
	store = int_v
	fmt.Println(store)
	store = string_v
	fmt.Println(store)

	// 并且可以使用 .(type) 的形式来判断是否存储了某类型的变量
	if v, ok := store.(int); ok {
		fmt.Println("store is (int)", v)
	} else if v, ok := store.(string); ok {
		fmt.Println("store is (string):", v)
	}
	store = int_v
	if v, ok := store.(int); ok {
		fmt.Println("store is (int)", v)
	} else if v, ok := store.(string); ok {
		fmt.Println("store is (string):", v)
	}

	// 或者使用 switch
	list := make(List, 3)
	list[0] = 1
	list[1] = "string"
	list[2] = Vertex{4, 7}

	for index, element := range list {
		switch value := element.(type) { // element.(type) 只能在 switch 里使用
		case int:
			fmt.Println(index, "(int)", value)
		case string:
			fmt.Println(index, "(string)", value)
		case Vertex:
			fmt.Println(index, "(Vertex)", value)
		}
	}

}
