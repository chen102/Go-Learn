package main

import "fmt"

func main() {
	a := 2
	if a == 3 { //左括号和if在同一行
		fmt.Println(a)
	} else {
		fmt.Println("not 3")
	}
	//if支持一个初始化语句,初始化和判断条件以分号分隔
	if s := 12; s == 10 {
		fmt.Println("s==10")
	} else if s > 10 { //多分枝
		fmt.Println("s>10")
	} else if s < 10 {
		fmt.Println("s<10")
	} else {
		fmt.Println("NO")
	}
	// fmt.Println(s)//undefined: s	s没有定义 在if初始化的变量，出if后就没用了
}
