//闭包通过匿名函数来实现
//所谓闭包就是一个函数捕获了和他在同一作用域的其他常量和变量
package main

import "fmt"

func main() {
	a := 10
	str := "hello"
	//匿名函数，没有函数名
	f1 := func() { //闭包，捕获同一个作用域外面的变量
		fmt.Println(a, str)
	}
	f1()

	//给一个函数类型起别名
	type functype func() //函数没有参数没有返回值
	//声明变量
	var f2 functype
	f2 = f1
	f2()

	//定义匿名函数，同时调用
	func() {
		fmt.Println(a, str)
	}() //后面的（）代表调用次匿名函数

	//带参数的匿名函数
	f3 := func(i, j int) {
		fmt.Println(i, j)
	}
	f3(3, 5)
	//同时调用
	func(i, j int) {
		fmt.Println(i, j)
	}(5, 3)

	//匿名函数有参有返回值
	x, y := func(i, j int) (max, min int) {
		if i > j {
			max = i
			max = j
		} else {
			max = j
			min = i
		}
		return
	}(4, 7)
	fmt.Println(x, y)

	//闭包捕获外部变量的特点
	func() {
		a = 666
		str = "no"
		fmt.Println(a, str)
	}()
	fmt.Println(a, str) //，外部也改了 即改的是同一个变量
	//闭包以引用方式捕获外部变量

	//传统局部变量
	fmt.Println(test())
	fmt.Println(test())
	fmt.Println(test())
	fmt.Println(test())

	//返回值为一个匿名函数，返回一个函数类型，通过f来调用返回的匿名函数，f来调用闭包函数
	f5 := test1()
	fmt.Println(f5())
	fmt.Println(f5())
	fmt.Println(f5())
	fmt.Println(f5())
	//他不关心这些捕获的变量和常量是否已经超出作用域
	//所以自要闭包还在使用它，这些变量就还会存在

}

//传统的局部变量
func test() int { //函数被调用时，x才分配空间，才初始化为0
	var x int
	x++
	return x * x //函数调用完毕，x自动释放
}

//函数的返回值是一个匿名函数，返回一个函数类型
func test1() func() int {
	var x int
	return func() int {
		x++
		return x * x
	}
}
