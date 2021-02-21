package main

import "fmt"

var o int

func test() {
	a := 10
	fmt.Println(a)
}
func test1() {
	fmt.Println(o)
}
func main() {
	//局部变量
	// a = 100 undefined: a
	{
		i := 10
		fmt.Println(i)
	}
	// i = 100 undefined: i
	if flag := 3; flag == 3 {
		fmt.Println(flag)
	}
	// flag = 100   undefined: flag
	//局部变量特点
	//定义在{}里面的变量就是局部变量，只能在{}里面有效
	//执行到定义变量那句话，才开始分配空间，离开作用域会自动释放
	//作用域，变量其作用的范围

	//全局变量
	o = 10
	fmt.Println(o)
	test1()
	//全局变量
	//定义在函数外部的变量是全局变量
	//全局变量可以在任何地方使用

	//如果局部比变量和全局变脸重名怎么办？ (不同作用域同名变量)
	var o byte
	fmt.Printf("%T\n", o)	//uint8
	{
		var o float64 
		fmt.Printf("%T\n",o)	//float64
	}
	//不同作用域，允许定义同名变量
	//使用变量的原则，就近原则
}
