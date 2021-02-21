package main

import "fmt"

func lost(x int) {
	result := 100 / x
	fmt.Println(result)
}
func main() {
	// //defer延时调用，main结束前调用
	// defer fmt.Println("aaaaaa")
	// fmt.Println("sssss")

	//多个defer执行顺序
	// fmt.Println("aaa")
	// fmt.Println("bbb")
	// //调用函数，导致内存出问题 panic：后面不会执行
	// lost(0)
	// fmt.Println("ccc")

	//多个defer语句，他会以	LIFO（先进先出），哪怕函数或是某个延时调用发生错误，这些调用依旧会执行
	defer fmt.Println("aaa")
	defer fmt.Println("bbb")
	defer lost(0)
	defer fmt.Println("ccc")
}
