package main

import "fmt"

//定义变量
func variableZeroValue() {
	var a int
	var s string
	fmt.Println(a, s)           //s字符串空打不出来
	fmt.Printf("%d %q\n", a, s) //println后面只能跟变量名，printf可以跟格式+变量名
}

//给变量赋值
func variableInitialValue() {
	//var c, b int = 3     //.\test.go:12:6: assignment mismatcha: 2 variable but 1 vlues
	var a, b int = 4, 5
	var s string = "abc"
	//fmt.Println(c, f) // command-line-arguments.\test.go:12:9: b declared and not used
	//go语言非常的严格 变量一旦定义了，就一定要用到 做项目是有效的避免定义一些无用的参数
	fmt.Println(a, b, s)
}

//int，string是多余的，编译器完全可以从后面的4，5 或是abc判断出来  让编译器自动决定类型
func variableTyoeDeduction() {
	var a, b, c, s = 4, 5, true, "def"
	fmt.Println(a, b, c, s)
}

//最简单的变量定义方法   第一次定义变量  var能不要就不用
func variableShorter() {
	a, b, c, s := 4, 5, true, "def"
	fmt.Println(a, b, c, s)
}

//在函数外定义变量 不能用上面的方法定义函数
var aa = 3 //在函数外部定义的变量不是“全局变量” （go语言没有全局变量的说法） 他是一个包内部的变量
var cc = 2

//一次定义多个变量
var (
	dd = 3
	ww = "cc"
)

func main() {
	fmt.Println("hello world")
	variableZeroValue()
	variableInitialValue()
	variableTyoeDeduction()
	variableTyoeDeduction()
	fmt.Println(aa, cc, dd, ww)
}
