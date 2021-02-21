package main

import (
	"fmt"
)

func test() (a, b, c int) { //go语言函数可以返回多个值
	return 1, 2, 3
}
func main() {
	// a, b, c := 10, 20, 30 //多重赋值

	// fmt.Println("a=", a, "b=", b, "c=", c) //println是一段一段处理的，自动换行
	// fmt.Printf("a=%d b=%d c=%d", a, b, c)  //格式化输出
	// a, b := 10, 20
	// fmt.Println(a, b)
	// a, b = b, a				//交换两个变量的值
	// fmt.Println(a, b)

	//tmp,_ i,j		//匿名变量，丢弃数据不处理，匿名变量配合函数返回值使用
	// var c, d, e int
	c, _, e := test() //不想要d这个变量
	fmt.Println(c, e) //1 0 3
}
