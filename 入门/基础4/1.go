package main

import "fmt"
import "math"

const test1 = "ohhhh"

func add(x, y int) int {
	return x + y
}
func swap(x, y string) (string, string) { //函数可以返回任意数量的返回值
	return y, x
}
func split(sum int) (x, y int) { //函数返回值可以被命名，没有参数的return语直接返回以命名的返回值
	x = sum + 1
	y = sum - 1
	return
}
func main() {
	fmt.Println(math.Pi) //只能应用其中以导出的名字（以大写字母开头）
	//如fmt,Println(math.pi)报错
	fmt.Println(add(33, 22))
	a, b := swap("hh", "tt")
	fmt.Println(a, b)
	var i int      //定义变量，go中定义的变量需要被使用
	fmt.Println(i) //没有赋值的int变量值为0,布尔为false，字符串为""
	//简洁赋值语句:=
	k := 3.2 //变量会根据初始值获得类型(int,float64,complex128)
	fmt.Printf("%T\n", k)
	//基本类型bool string int(|8|16|32|64) uint(|8)|16|32|64) byte=uint8 rune=int32 float32 float64 complex64 complex128
	//常量使用const关键字
	const test = "hello" //常量可以是字符，字符串，布尔值或数值。定义的常量不需要一定用
	fmt.Println(test)
	fmt.Println(test1)
	//iota 第一个iota等于0，每当iota在新的一行被使用时，他的值会自动加1
	const (
		l = iota //0
		c        //1
		d = 1    //1
		j        //1
		e = iota //恢复计数 3
		o        //4
	)
	fmt.Println(l, c, d, j, e, o)
	//iota还有一个用法
	const (
		no1 = 1 << iota
		no2 = 3 << iota //3左移一位 110即6
		no3             //3左移一位 1100 即12
		no4             //3左移一位 11000 即24
	)
	fmt.Println(no1, no2, no3, no4)
	const (
		cc1 = 1 << iota
		cc2
		cc3
		cc4
	)
	fmt.Println(cc1, cc2, cc3, cc4)

}
