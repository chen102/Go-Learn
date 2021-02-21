package main

import "fmt"

func main() {
	var p *int
	p = new(int) //动态分配空间
	*p = 10
	fmt.Println(*p)
}
