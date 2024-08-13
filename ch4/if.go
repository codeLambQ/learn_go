package main

import "fmt"

// if bool 表达式 { 处理语句 }

func main() {
	age := 16
	country := "中国"
	if age < 18 && country == "中国" { // 也可以使用多种条件
		fmt.Println("未成年")
	} else if age == 18 {
		fmt.Println("刚刚成年")
	} else {
		fmt.Println("已成年")
	}

	// 也可以这样写

	if age < 18 {
		fmt.Println("未成年")
	}

	if age == 18 {
		fmt.Println("刚刚成年")
	}

	if age > 18 {
		fmt.Println("已成年")
	}
}
