package main

import "fmt"

// map 就是 Python 中 dict 的概念
// 声明 var variable map[key-type] value-type

func main() {
	// var fake_numbers map[string]int
	// fake_numbers["one"] = 1
	// 上面的语句会报错：assignment to entry in nil map
	// 在使用 map 之前需要 make
	var numbers map[string]int = make(map[string]int)
	numbers["one"] = 1
	numbers["two"] = 2
	fmt.Println(numbers["one"], numbers["two"]) // 1, 2

	// Map 是无序的
	MapIsDisorder := make(map[string]int)
	MapIsDisorder["one"] = 1
	MapIsDisorder["two"] = 2
	MapIsDisorder["three"] = 3
	MapIsDisorder["four"] = 4
	MapIsDisorder["five"] = 5
	fmt.Println("MapIsDisorder =", MapIsDisorder)           // MapIsDisorder = map[one:1 two:2 three:3 four:4 five:5]
	fmt.Println("len(MapIsDisorder) =", len(MapIsDisorder)) // len(MapIsDisorder) = 5

	// range: Go 的 range 和 Python 的 range 不一样= =
	// 遍历 MapIsDisorder 输出 *2 的结果
	for k, v := range MapIsDisorder {
		fmt.Println(k, " => ", v)
	}
	// one => 1
	// two => 2
	// three => 3
	// four => 4
	// five => 5

	// Map 初始化
	OfficialURL := map[string]string{"Go": "http://golang.org/", "Python": "http://www.python.org/"}
	fmt.Println(OfficialURL) // map[Go:http://golang.org Python:http://www.python.org]

	// Map 获取值的时候返回两个值，第二个值为一个 bool 值，当 key 存在的时候返回 true，否则返回 false
	value, ok := OfficialURL["Ruby"]
	val := OfficialURL["Ruby"]                  // == val, _ := OfficialURL["Ruby"]
	fmt.Println("value =", value, "\nok =", ok) // value=  ok = false
	fmt.Println("val =", val)                   // val =

	// 添加、修改、删除
	OfficialURL["Ruby"] = "https://www.ruby-lang/en/"
	fmt.Println(OfficialURL["Ruby"]) // https://www.ruby-lang/en/
	OfficialURL["Ruby"] = "https://www.ruby-lang.org/en/"
	fmt.Println(OfficialURL["Ruby"])         // https://www.ruby-lang.org/en/
	fmt.Println("length:", len(OfficialURL)) // length: 3
	delete(OfficialURL, "Ruby")
	fmt.Println("length:", len(OfficialURL)) // length: 2

}
