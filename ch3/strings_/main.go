package main

import (
	"fmt"
	"strings"
)

func main() {
	// strings 的一些方法
	name := "imooc体系课-go工程师"
	fmt.Println(name)

	// 查看是否包含
	fmt.Println(strings.Contains(name, "go"))

	// 查看 go 出现的位置，现在还是将中文算成两个位置
	fmt.Println(strings.Index(name, "go"))

	// 计算字符串中某个字符串线了多少次
	fmt.Println(strings.Count(name, "go"))

	// 替换,方法左右一个值小于0的时候替换字符串中所有符合的，大于0的时候只会替换第一个
	fmt.Println(strings.Replace(name, "go", "java", -1))
}
