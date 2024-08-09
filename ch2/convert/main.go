package main

import (
	"fmt"
	"strconv"
)

func main() {
	// 类型转换
	// int 转 uint
	var a int = 12
	b := uint8(a)

	fmt.Println(b)

	// 定义个别名
	type IT int
	//var c IT = int(a) // 这种是不被允许的
	var c IT = IT(a) // 这种是正确写法
	fmt.Println(c)

	// 字符串转数子
	var s = "123"
	myint, err := strconv.Atoi(s) // 会有错误
	if err != nil {
		fmt.Println("convert error")
	}
	fmt.Println(myint)

	// 数字转字符串
	var iros = 33
	fmt.Println(strconv.Itoa(iros)) // 数字转字符串一般不会出现问题，所以可以直接转

	// 字符串转换成为基本类型  ParseXXX()
	// 字符串转换成浮点类型
	float, err := strconv.ParseFloat("3.45", 64)
	if err != nil {
		return
	}
	fmt.Println(float)

	parseInt, err := strconv.ParseInt("123", 10, 64)
	if err != nil {
		return
	}
	fmt.Println(parseInt)

	// 基本类型转换成字符串  format
	boolstr := strconv.FormatBool(true)
	fmt.Println(boolstr)

	float22 := strconv.FormatFloat(2.222222, 'E', -1, 64)
	fmt.Println(float22)

	strint := strconv.FormatInt(23123, 10)
	fmt.Println(strint)
}
