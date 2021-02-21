package main

import "fmt"
import "math"

//定义常量
func consts() {
	const (
		filename = "test.txt"
		a, b     = 3, 4 //可以规定类型，也可以不规定,不规定类型，常量的类型是不确定的
	)
	var c int
	c = int(math.Sqrt(a*a + b*b)) //传进Sqrt的参数（a*a + b*b）不用强制转换成float了，因为常量a，b已经当float来用了
	fmt.Println(filename, c)
}

//const数值可作为各种类型使用

//特殊的常量类型-枚举型
//普通枚举型
//go语言没有特殊的枚举关键字，一般用一组const定义枚举

func enums() {
	const (
		a = 0
		b = 1
		c = 2
		d = 3
	)
	fmt.Println(a, b, c, d)
}

//go语言为一组const做了简化
//自增枚举型
func enums1() {
	const (
		a = iota //元素iota:这组const是自增值   默认从零开始
		f
		c
		d
	)
	//iota也能参与运算即可以作为一个自增值的种子
	const (
		b = 1 << (10 * iota)
		kb
		mb
		gb
		tb
		pb
	)
	fmt.Println(a, c, f, d)
	fmt.Println(b, kb, mb, gb, tb, pb)
}
func main() {
	consts()
	enums()
	enums1()
}
