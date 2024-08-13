package main

import (
	"fmt"
	"time"
)

func main() {
	// for循环，在 go 语言中没有 while 语言，可以使用 for 来代替 while

	for i := 0; i < 4; i++ {
		fmt.Println(i)
	}

	fmt.Println("---------------")

	// 另一种写法
	var i int
	for ; i < 3; i++ {
		fmt.Println(i)
	}

	fmt.Println("---------------")

	var i2 int = 0
	for i2 < 3 {
		fmt.Println(i2)
		i2++
	}

	fmt.Println("---------------")
	// 死循环
	var i3 int
	for {
		time.Sleep(10) // 休眠10ms
		fmt.Println(i3)
		i3 += 1
		if i3 == 3 {
			break
		}
	}

	fmt.Println("---------------, 1加到100的值")

	// 1+100
	var sum int
	for i := 1; i <= 100; i++ {
		sum += i
	}
	fmt.Println(sum)

	fmt.Println("---------------, 打印乘法口诀")
	for i := 1; i <= 9; i++ {
		for j := 1; j <= i; j++ {
			c := i * j
			fmt.Printf("%d*%d=%d ", i, j, c)
		}
		fmt.Println()
	}

	fmt.Println("---------------,for range, 类似于 java 的 for each")
	// 主要对字符串、数组、切片、map、channel 使用

	name := "imooc go"
	for index, value := range name {
		fmt.Println(index, value)
		fmt.Printf("%d %c\r\n", index, value)
	}
}
