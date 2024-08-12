package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	// 长度计算
	name := "immocsda"
	fmt.Println(len(name)) // 如果只有英文的话可以直接使用 len 方法计算，如果有中文的话不能这样计算

	name = "go体系课"
	fmt.Println(len(name)) // 这里打印出来的结果是 11，显然是不对的，因该打印是5
	// 可以使用下面这种方法
	bytes := []rune(name)
	fmt.Println(len(bytes)) // 这样输出结果就是 5 了

	// 转义字符，让改变当前字符的意思，使用 \ 来对字符进行转移
	// newName := "ll"lulu"" // 如果想在字符串里直接打印出来双引号是不行的，需要使用 \ 进行转移
	newName := "ll\"lulu\""
	fmt.Println(newName)
	// 或者
	newName = `ll"lulu"`
	fmt.Println(newName)

	// 格式化输出
	name = "lulu"
	age := 26
	adress := "beijing"

	fmt.Printf("姓名: %s 年龄: %d 地址: %s\n", name, age, adress)
	// 也可以使用有返回值的格式化输出
	strp := fmt.Sprintf("姓名: %s 年龄: %d 地址: %s", name, age, adress)
	fmt.Println(strp)

	// 高性能字符串拼接方法，builder
	var builder strings.Builder
	builder.WriteString("姓名: ")
	builder.WriteString(name)
	builder.WriteString("年龄: ")
	builder.WriteString(strconv.Itoa(age))
	builder.WriteString("地址: ")
	builder.WriteString(adress)
	re := builder.String() // 将写入进 builder 里面的内容赋值给 re
	fmt.Println(re)
}
