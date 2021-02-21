package main

import "fmt"

func main() {
	s := "fasfadgf"
	a := make(map[byte]int)
	for _, v := range []byte(s) {
		fmt.Println(a[v])
	}

}
