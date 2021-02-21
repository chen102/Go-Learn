//写一个去头去尾的函数
package main

import "fmt"

func test(s []int) int {
	s = s[1:]
	s = s[:len(s)-1]
	return s
}
func main() {
	s := []int{1, 3, 4, 5, 6}
	test(s[:])
	fmt.Println(s)
}
