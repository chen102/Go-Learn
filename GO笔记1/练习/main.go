package main

import (
	"fmt"
)

func sx(a, b int) (int, int) {
	if a > b {
		return a, b
	} else {
		a, b = b, a
		return a, b
	}

}
func main() {
	fmt.Println(sx(9, 4))
}
