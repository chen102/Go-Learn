package main

import "fmt"

func main() {
	a := 10
	str := "hello"

	defer func(a int, str string) {
		fmt.Println("defer后", a, str)
	}(a, str) //10 hello 把参数传递过去，已经先传递参数，只是因为defer没有调用

	a = 100
	str = "no"
	fmt.Println("外部", a, str)

	func(a int, str string) {
		fmt.Println("没有defer", a, str)
	}(a, str) //100 no
}
