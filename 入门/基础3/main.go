package main

import "fmt"

func test(arg ...int) {
	for _, s := range arg {
		fmt.Println(s)
	}
}
func test1(arg ...int) {
	a := len(arg) - 2
	test(arg[a:]...)
	//for _, s := range arg {
	//		fmt.Println(s)
	//	}
}
func test2(a, b int, s func(int, int) int) int {
	return s(a, b)
}
func add(a, b int) int {
	return a + b
}

type my_functype func(int, int) int //函数别名

func test3(a, b int, s my_functype) int {
	return s(a, b)
}
func test4() func() int { //返回一个匿名函数，及一个函数类型
	var x int
	return func() int {
		x++
		return x
	}
}
func main() {
	test1(1, 2, 3, 4, 5, 6, 7)
	fmt.Println(test2(2, 4, func(a, b int) int { //匿名函数
		return a - b
	}))
	fmt.Println(test2(2, 4, add))
	fmt.Println(test3(7, 4, add))
	var test my_functype //通过变量来调用函数
	test = add
	fmt.Println(test(6, 3))

	s := func(i, j int) (s int) {
		s = i + j
		return
	}(4, 2)
	fmt.Println(s)
	f := test4() //返回是一个函数类型，通过变量来调用匿名函数
	fmt.Println(f())
	fmt.Println(f())
	fmt.Println(f())
	fmt.Println(f())
	arr := [...]int{1, 2, 3, 4, 5, 6}
	arr1 := arr[:2]
	arr2 := append(arr1, arr[2:]...)
	fmt.Println(arr2)
}
