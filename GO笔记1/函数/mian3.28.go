package main

import (
	"fmt"
)

//简单计算器
func eval(a, b int, str string) (int, error) { //返回值的类型写在最后面，可返回多个值，函数作为参数，没有默认参数，可选参数，只有一个可变参数列表
	switch str {
	case "+":
		return a + b, nil
	case "-":
		return a - b, nil
	case "*":
		return a * b, nil
	case "/":
		return a / b, nil
	default:
		return 0, fmt.Errorf("unsupported operation: %s", str)
	}
}
func div(a, b int) (int, int) { //多个返回值
	return a / b, a % b
}
func divs(a, b int) (q, r int) { //给返回值一个名字
	q = a / b
	r = a % b
	return
}

//函数式编程
// func apply(str func(int, int) int, a, b int) int {
// 	p := reflect.ValueOf(str).Pointer()
// 	runtime.FuncForPC(p).Name() //获取函数名
// 	fmt.Printf("Callint function %s with args "+"(%d,%d)", strName, a, b)
// 	return str(a, b)
// }
func father(son func(int, int) int, a, b int) int {
	return son(a, b)
}

// func son(a, b int) int {
// 	return a + b
// }

//函数可变参数列表
func sum(num ...int) int { //...int:任意几个
	s := 0
	for i := range num {
		s += num[i]
	}
	return s
}

func main() {

	// if result, err := eval(5, 2, "*"); err != nil {
	// 	fmt.Println("Error", err)
	// } else {
	// 	fmt.Println(result)
	// }

	// fmt.Println(div(11, 2))
	// c, d := divs(11, 2)
	// fmt.Println("c =", c, "d = ", d)

	// fmt.Println(father(func(a, b int) int {
	// 	return a + b
	// }, 5, 8)) 								//函数的参数可以作为函数---函数式编程

	fmt.Println(sum(1, 2, 3, 4, 5, 55, 7)) //求和
}
