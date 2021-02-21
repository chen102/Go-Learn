//package main

import "fmt"

func main() {
	for i := 0; i < 10; i++ {
		fmt.Println(i)
	}
	i := 10
	for i > 0 { //go中没有while
		i--
		fmt.Println(i)
	}
	v := 10
	for { //死循环
		if v <= 0 {
			break
		} else {
			v--
		}
	}
	fmt.Println("OK")

}
