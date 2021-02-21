package main

import (
	"fmt"
	"io/ioutil"
)

//if判断语句，if的条件里不需要括号
func ifs() {
	const filename = "a.txt"
	con, err := ioutil.ReadFile(filename) //库函数ioutil.ReadFile:将文件中内容读出来，函数会返回两个值（go语言的函数可以返回两个值）
	if err != nil {                       //nic:ioutil.ReadFile函数返回的参数
		fmt.Println(err) //open a.txt: The system cannot find the file specified.  在工作目录下没有a.txt
	} else {
		fmt.Printf("%s\n", con) //读出文件内容
	}
}

//另一种写法,if的条件里可以赋值，if的条件里赋值的变量作用域就在这个if语句里
func ifss() {
	const filename = "a.txt"
	if con, err := ioutil.ReadFile(filename); err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("%s\n", con)
	} //这里的con，err是在if里定义的，出if后就失效了
	// fmt.Printf("%s\n", con)          .\main3.25.go:27:21: undefined: con
}

//switch判断语句，switch会自动break，除非使用fallthrough
func grade(score int) string { //传入score,传出string
	g := ""
	switch {
	case score < 0 || score > 100:
		panic(fmt.Sprintf("Wrong score: %d", score)) //panic：中断程序执行,下面的语句都不会执行
	case score < 60:
		g = "F" //switch后也可以没有表达式 如：return "F"
	case score < 80:
		g = "C"
	case score < 90:
		g = "B"
	case score <= 100:
		g = "A"
	}
	return g
}

func main() {
	//ifs()
	//ifss()
	fmt.Println(
		grade(0),
		grade(59),
		grade(60),
		grade(99))
	// grade(101)   // grade（-3）
}
