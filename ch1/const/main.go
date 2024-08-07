package main

import "fmt"

func main() {
	// 常量，使用 const 关键字进行定义，定义之后不可以再更改了，变量名一般大写
	const PI float32 = 3.1415926 // 定义变量类型为显示定义，现定义类型属于隐式定义
	// const PI = 3.1415926

	// 可以使用组来进行定义
	const (
		UNKNOWN = 1
		FEMALE  = 2
		MALE    = 3
	)

	const (
		x int = 16
		y
		z string = "lulu"
		s
	)

	fmt.Println(x, y, z, s)

	/*
		常量类型只可以定义 bool 数值类型(整形，浮点型) 和字符串
		常量等号左右的类型需要一致
		常量定义后不一定非要使用
		常量分组定义后，如果不给变量赋值的话会默认和上一个常量的类型数值一样
	*/
}
