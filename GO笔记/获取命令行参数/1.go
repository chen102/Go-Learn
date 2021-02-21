package main

import (
	"fmt"
	"os"
)

func main() {
	list := os.Args
	n := len(list)
	fmt.Println(n)
	for i, v := range list {
		fmt.Println(i, v)
	}
}
