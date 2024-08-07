package main // 第一行必须为 package main 不然不可以使用
import "fmt"

// 全局变量
// var name = "lulu"
// var gae = 2
// 可以将上面进行合并
var (
	name = "lulu"
	age  = 2
)

// 下面这种方式不可以用来定义全局变量
// adress := "sad"

// 全局变量定义之后可以不使用，局部变量不可以不使用，不使用的话会报错
func main() {
	// go 变量的定义，分为全局变量和局部变量
	// 局部变量
	// go 语言定义变量指定类型的位置非常的奇怪，在变量名称的后面，有别于其它的静态语言
	var age int
	age = 1

	// 也可以将两行合并成一盒
	var name = "lulu"

	// 更简易的定义方式
	address := "beijing"

	// 多变量的定义方式
	var name1, name2, name3 = "sd", "wsda", 4

	fmt.Println(age, name, address)

	fmt.Println(name1, name2, name3)

	/*
		变量总结
		1.变量必须先定义在使用
		2. 变量的类型确定后不可以赋值不同类型的值
		3.变量名不能冲突
		4.变量有默认值
		5. 全局变量定以后可以不实用，局部变量不可以不使用，编译会报错
	*/

}
