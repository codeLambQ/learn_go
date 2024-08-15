package main

import "fmt"

// 数组
func main() {
	var str [3]string // str 只有三个元素的数组
	// []string 是切片和 [3]string 不是同一个类型

	// 给数组赋值
	str[0] = "go"
	str[1] = "grpc"
	str[2] = "gin"

	// 遍历数组
	for _, value := range str {
		fmt.Println(value)
	}

	fmt.Println("数组的初始化1")
	arr1 := [3]string{"go", "grpc", "gin"}
	for _, value := range arr1 {
		fmt.Println(value)
	}

	fmt.Println("数组的初始化2,在指定位置上进行初始化")
	arr2 := [3]string{1: "grpc"}
	for _, value := range arr2 {
		fmt.Println(value)
	}

	fmt.Println("数组的初始化3，不固定元素个数的数组")
	arr3 := [...]string{"go", "grpc", "gin"}
	fmt.Printf("%T\n", arr3)
	for _, value := range arr3 {
		fmt.Println(value)
	}

	fmt.Println("数组可以直接使用==来进行比较，如果相同位置的元素一样则认为数组相等")
	if arr1 == arr3 {
		fmt.Println("arr1 == arr3")
	}

	// 多维数组
	fmt.Println("多维数组的创建、赋值和打印")
	arrInfo := [3][4]string{} // 定义一个3行4列的数组
	arrInfo[0] = [4]string{"go", "grpc", "gin", "入门"}
	arrInfo[1] = [4]string{"go", "grpc", "gin", "入门"}
	arrInfo[2] = [4]string{"go", "grpc", "gin", "入门"}

	for i := 0; i < len(arrInfo); i++ {
		for j := 0; j < len(arrInfo[i]); j++ {
			fmt.Print(arrInfo[i][j] + " ")
		}
		fmt.Println()
	}

	// 使用 forrange
	for _, value := range arrInfo {
		for _, info := range value {
			fmt.Print(info + " ")
		}
		fmt.Println()
	}
}
