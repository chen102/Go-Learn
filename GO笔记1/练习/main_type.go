package main

import "fmt"

type myfunc func(float64, float64, string) float64 //定义一个函数类型myfunc		func后面的数据类型必须要与被调用的函数函数一致 同样是两个int参数一个int返回值
func clac(a, b float64, s string, funcs myfunc) (resoult float64) {
	resoult = funcs(a, b, s)
	return
}
func jjcc(a, b float64, s string) float64 {
	switch {
	case s == "cheng":
		return a * b
	case s == "chu":
		return a / b
	case s == "jia":
		return a + b
	case s == "jian":
		return a - b
	default:
		fmt.Println("Please input jia,jian,cheng,chu")
	}
}

// func jia(a, b float64) float64 {
// 	return a + b
// }
// func jian(a, b float64) float64 {
// 	return a - b
// }
// func cheng(a, b float64) float64 {
// 	return a * b
// }

// func chu(a, b float64) float64 {
// 	return a / b
// }

func main() {
	so := clac(5, 3, "cheng", jjcc) //funcs不用定义
	fmt.Printf("%v", so)
}
