package main

import "fmt"

func main() {
	a := 10
	fmt.Println(a)  //变量的内存
	fmt.Println(&a) //变量的地址

	var p *int = &a
	*p = 5 //操作的是p所指向的内存地址
	fmt.Println(a, *p)
	fmt.Println(&a, p)

	var pa *int
	fmt.Println(pa) //没有指向默认是<nil>

	c, d := 2, 6
	swap(&c, &d)
	fmt.Println(c, d)

}
func swap(a, b *int) {
	*a, *b = *b, *a
}
