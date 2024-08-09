package main

func main() {
	// go 数据类型 整数型
	//var a int8    有符号 8位
	//var b int16	有符号 16位
	//var c int32	有符号 32位
	//var d int64	有符号 64位
	//var ua uint8	无符号 8位
	//var ub uint16	无符号 16位
	//var uc uint32	无符号 32位
	//var ud uint64	无符号 64位
	//var i int  动态类型，根据你的计算机来决定如果是32位就是 int32 如果是 32 位 就是 64 位

	//var b byte    字符类型，go语言没有提供字符类型，因为字符的本质也是数字，直接使用byte进行表示，byte 其实就是 uint8,
	// 在 go 与严重叫做别名，源码 type byte = uint8
	// var b byte = 'a',这样直接输出，谁输出数字，如果想要输出成字符串，则需要进行格式化输出，fmt.Printf("b=%c", b)

	//var r rune  也是字符类型，本质上是 int32, 用来表示非英文的字符比较合适，如果想要输出成中文过着其他语言，则需要格式化输出

	//var s string = "ssdassdaf" 字符串类型，比较简单，写入啥，输出啥

	//var f1 float32  浮点类型，必须指定位数
	//var f2 float64

}
