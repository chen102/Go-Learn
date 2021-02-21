package main

import "fmt"

//无参无返回值函数的定义
func test() {
	a := 10
	fmt.Println(a)
}

//有参无返回值函数的定义，普通参数列表
func test1(a, b int) { //定义函数时，在函数名后面（）定义的参数叫形参

	fmt.Println(a, b)
}

//不定参数列表

func test3(args ...int) { //...int:...type不定参数类型，传递的实参可以是0个或多个
	fmt.Println(len(args))

	fmt.Println(args) //获取用户传递的参数
	for i := 0; i < len(args); i++ {
		fmt.Println(args[i])
	}
	//或者用迭代
	for i, v := range args {
		fmt.Printf("args[%v] = %v \n", i, v)
	}
	//或者用切片
	fmt.Println(args[:])
}

//不定参数，一定（只能）放在形参中的最后一个参数
// func test5(args ...int ,a int){		cannot use ... with non-final parameter args
// 	fmt.Println(a, args)
// }
func test4(a int, args ...int) { //固定参数一定要传
	fmt.Println(a, args)
}

//不定参数传递
func test6(args ...int) { //将参数传个另一个函数用
	fmt.Println(args)
}
func test7(args1 ...int) {
	//test6(args1...) //函数里调函数,全部元素传递给test6
	//永远只传后面两位元素给test6
	var a = len(args1) - 2
	test6(args1[a:]...)
}

//无参有返回值
//一个返回值，有返回值的函数需要通过return中断函数，通过return返回
func test8() int {
	return 666
}

//给返回值起一个变量名，go推荐写法
func test9() (result int) {
	result = 666
	return //默认返回result
}

//多个返回值
func test10() (a, b int) {
	a, b = 3, 4
	return
}

//有参有返回值
func test11(a, b int) (max, min int) {
	if a > b {
		max, min = a, b
	} else {
		max, min = b, a
	}
	return
}
func main() {
	test()
	test1(666, 999) //参数传递，只能由实参传递给形参，不能反过来
	test3(4, 3, 2, 1)
	test4(1, 2, 3, 4, 5, 6)
	test7(1, 2, 3, 4, 5, 6)
	fmt.Println(test8())
	fmt.Println(test9())
	fmt.Println(test10())
	max, min := test11(1, 4)
	fmt.Printf("max=%d min=%d\n", max, min)
	a, _ := test11(1, 4)
	fmt.Printf("max=%d", a)
}
