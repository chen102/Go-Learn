package main

import "fmt"

func main() {
	a, b, c, d := 10, "abc", 'a', 3.14
	fmt.Printf("%T,%T,%T,%T\n", a, b, c, d)
	fmt.Printf("%d,%s,%c,%f\n", a, b, c, d)
	fmt.Printf("%v,%v,%v,%v\n", a, b, c, d) //%v自动匹配格式
	//%s字符串
	//%c字符
	//%d整形
	//%T浮点型
}
