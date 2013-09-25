package main

import "fmt"

func main() {

	// bool 类型默认值是 false
	var thisWillBeFalse bool
	valid := true // 简短声明
	fmt.Println(thisWillBeFalse, valid)

	// int
	var int_v int = -1
	fmt.Println(int_v)

	// unsigned int
	var uint_v uint = 2
	fmt.Println(uint_v)

	// int32
	var int32_v int32 = 3
	fmt.Println(int32_v)

	// alias of int32
	var rune_v rune = 3
	fmt.Println(rune_v)

	// Also: int8, int16, int32, int64, uint8, uint16, uint32, uint64
	// byte is alias of uint8

	// float32, float64
	var float32_v float32 = 2.5
	fmt.Println(float32_v)

	// complex64, complex128 (实数、虚数各 64 位)
	var complex128_v complex128 = 5 + 5i
	fmt.Println(complex128_v)

	// string
	var string_v string = "A string"
	var str_v string = `A String too`
	fmt.Println(string_v, str_v)

	// 字符串是不能修改的，所以 str[0] = 'x' 是错误的，但是你可以将字符串转换成 []byte
	str_v_byte := []byte(str_v)
	str_v_byte[2] = 's'
	str_v_str := string(str_v_byte)
	fmt.Println(str_v_str)

	// 字符串连接使用 '+'
	str_join := string_v + ", " + str_v
	fmt.Println(str_join)

	// 切片
	str_slice := string_v[2:]
	fmt.Println(str_slice)

	// 定义多行字符串使用 '`'
	str_multi_line := `Mu~~~~~
    lti-line string~~`
	fmt.Println(str_multi_line)

	// nil
	fmt.Println(nil)
}
