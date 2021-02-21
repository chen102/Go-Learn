// hello
package main

import (
	"fmt"
	"math"
)

func vars() {
	a := 10
	var b = 20
	const f, g = 3, 4
	fmt.Println(math.Sqrt(f*f + g*g))
	fmt.Println(a + a + b + f)
}

// func ifs(a) {
// var a int
// if a >= 0 {
// fmt.Printf("正")
// } else {
// fmt.Printf("负")
// }

// }
func main() {
	fmt.Println("Hello World!")
	vars()
	// ifs(5)
}
