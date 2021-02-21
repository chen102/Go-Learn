package main

import "fmt"

func main() {
	jjcc(5, 3, cheng)
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
