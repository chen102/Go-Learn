package main

import "fmt"

func main() {
	// var (
	// 	a = 10
	// 	b = 2.5
	// )
	// c, d := 2.2, 4
	// fmt.Println(a, b, c, d)
	// const (
	// 	f = 20
	// 	g = 30
	// )
	// fmt.Println(f, g)
	//iota常量自动生成器，每隔一行，自动累加1
	const (
		a = iota //iota遇到const，重置为零
		b
		c
		d
	)
	fmt.Println(a, b, c, d)
}
