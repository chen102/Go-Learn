package main

import "fmt"

func main() {
	arr := [...]int{0, 1, 2, 3, 4, 5, 6}
	fmt.Println(len(arr), cap(arr))
	s := arr[:3]
	fmt.Println(len(s), cap(s), s)
	s1 := s[:4] //Slice可扩展
	fmt.Println(s1)
	const n = 30
	// fmt.Println("多少：")
	// fmt.Scan(&n)
	var arr1 [n]int //数组下标不能用函数？
	for i := 0; i <= n-1; i++ {
		arr1[i] = i
	}
	fmt.Println(arr1)
	arr2(arr[:3]) //切片
}
func arr2(arr []int) {
	for _, v := range arr {
		fmt.Println(v)
	}

}
