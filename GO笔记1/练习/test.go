package main

import "fmt"

func father(son func(int, int) int, a, b int) int {
	return son(a, b)
}

// func son(a, b int) int {
// 	return a + b
// }

func main()
 {
	fmt.Println(father(func(a, b int) int {
		return a + b
	}, 5, 8)) //函数的参数可以作为函数---函数式编程
}
// func main() {
// 	fmt.Println(ff(func(a, b int) int {
// 		return a + b
// 	}, 5, 8))
// }
