package main

import "fmt"

func main() {
	s := "fasfasf"
	for i, v := range s { //range:迭代 会返回两个值，第一个返回值是元素的数组下标，第二个返回值是元素的值
		// fmt.Println(i, v) //默认返回的是int型
		fmt.Printf("%d %c\n", i, v)
	}
	for _, r := range s {
		fmt.Printf("%c\n", r) //用匿名变量的方式只查看元素的值
	}
	arr1 := [5]int{3, 75, 34, 6, 4} //数组的便利
	for q, w := range arr1 {
		fmt.Println(q, w)
	}
}
