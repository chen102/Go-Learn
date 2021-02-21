//函数调用流程：先调用后返回，先进后出
package main

import "fmt"

func funcb() (b int) {
	b = 222
	fmt.Println("funcb b = ", b)
	return
}
func funca() (a int) {
	a = 111
	b := funcb()
	fmt.Println("funca b = ", b)
	fmt.Println("funca a = ", a)
	return
}

//函数递归，调用自己
func test(a int) {
	if a == 1 {
		fmt.Println("a =", a)
		return //终止函数调用
	}
	test(a - 1)
	fmt.Println("a =", a)
}

//递归实现1+2+3+.....100
func test1(i int) int {
	if i == 1 {
		return 1
	}
	return i + test1(i-1)
}
func test2(i int) int {
	if i == 100 {
		return 100
	}
	return i + test2(i+1)
}

//函数类型
func add(a, b int) int {
	return a + b
}
func minus(a, b int) int {
	return a - b
}

//函数也是一种数据类型，通过type给一个函数类型起名
type functype func(int, int) int //数据类型要和上面的一致，有相同的参数，相同的返回值

//回调函数
//函数有一个参数是函数类型，这个函数就是回调函数
//函数式编程？？？
//多态，多种形态，调用同一个接口，不同的表现。可以实现不同表现
//先有想法，后面再实现功能
func calc(a, b int, ftest functype) (result int) { //函数类型必须放后面
	result = ftest(a, b)
	return
}
func main() {
	a := funca()
	fmt.Println("main a = ", a)
	test(3)
	fmt.Println("main")
	sum := test1(100)
	fmt.Println(sum)
	sum1 := test2(1)
	fmt.Println(sum1)

	// result := add(1, 3) //传统的调用方式
	// fmt.Println(result)
	// var ftest functype //通过变量来调用函数
	// ftest = add
	// result1 := ftest(3, 4) //等价于add(3,4)
	// fmt.Println(result1)

	// ftest = minus
	// result2 := ftest(7, 4) //等价于minus(3,4)
	// fmt.Println(result2)

	a1 := calc(3, 4, add)
	a2 := calc(6, 3, minus)
	fmt.Println(a1)
	fmt.Println(a2)

}
