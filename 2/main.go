package main

import (
	"fmt"
	"io/ioutil"
	"math"
	"reflect"
	"runtime"
)

//强制类型转换
func test() {
	a, b := 3, 4
	var c int
	c = int(math.Sqrt(float64(a*a + b*b)))
	fmt.Println(c)
}
func iota1() {
	const (
		b = 1 << (10 * iota)
		kb
		mk
		gb
		tb
		pb
	)
	fmt.Println(b, kb, gb, tb, pb)
}
func test2() {
	const filename = "abc.txt"
	if contents, err := ioutil.ReadFile(filename); err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("%s\n", contents)
	}

}

//函数式编程
func apply(op func(int, int) int, a, b int) int {
	p := reflect.ValueOf(op).Pointer()
	opName := runtime.FuncForPC(p).Name()
	fmt.Println("Calling function %s with args "+"(%d,%d)", opName, a, b)
	return op(a, b)
}

//可变参数列表
func sum(n ...int) int {
	s := 0
	fmt.Printf("%T", n)
	for i := range n {
		s += n[i]
	}
	return s
}
func test3() {
	a := 10
	p := &a
	*p = 11
	fmt.Println(a, *p)
}
func main() {
	test()
	iota1()
	test2()
	fmt.Println(apply(
		func(a, b int) int { //匿名函数
			return int(math.Pow(
				float64(a), float64(b)))
		}, 3, 4))
	fmt.Println(sum(3, 35, 6, 3))
	test3()
}
