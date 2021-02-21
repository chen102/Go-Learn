package main

import "fmt"

func main() {
	// flag := true
	// fmt.Printf("flag = %t", flag)
	//布尔型转整形
	// fmt.Printf("flag = %d", int(flag))   //.\main.go:9:29: cannot convert flag (type bool) to type int
	//布尔类型不能转换为整形

	//整形转换为布尔类型
	// flag1 := 1
	// fmt.Printf("flag = %t", bool(flag1)) //.\main.go:14:30: cannot convert flag1 (type int) to type bool
	//整形也不能转化为布尔型

	//这种不能转换的类型，叫不兼容类型

	// var ch byte
	// ch = 'a'
	// var t int
	// t = int(ch)
	// fmt.Println(t)

	//类型别名
	type bigint int64 //给int64起一个别名bigint
	var a bigint
	fmt.Printf("%T", a)

}
