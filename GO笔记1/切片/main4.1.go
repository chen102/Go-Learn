//Slice(切片)
package main

import "fmt"

func main() {
	arr := [...]int{0, 1, 2, 3, 4, 5, 6}
	s := arr[2:6] //切数组第二个元素到第六个元素中间的值  半开半闭区间（2包括进去，6不包括）
	fmt.Println(s)

	fmt.Println(arr[2:6], arr[:6], arr[2:], arr[:])

	//Slice是对数组的一个视图
	//Slice本身没有数据，是对底层array的一个view
	s1 := arr[2:]
	s2 := arr[:]
	fmt.Println(s1) //[2 3 4 5 6]
	fmt.Println(s2) //[0 1 2 3 4 5 6]

	updateSlice(s1)

	fmt.Println("After updateSlice ")
	fmt.Println(s1) //[100 3 4 5 6]
	fmt.Println(s2) //[0 1 100 3 4 5 6]

	updateSlice(s2)

	fmt.Println("After updateSlice ")
	fmt.Println(s1) //[100 3 4 5 6]
	fmt.Println(s2) //[100 1 100 3 4 5 6]

	fmt.Println(arr) //[100 1 100 3 4 5 6]
	//arr也被改掉了

	arr2 := [...]int{0, 1, 2, 3, 4, 5}
	printArray1(arr2[:])
	fmt.Println(arr2)

	//Reslice
	s3 := arr2[:]
	fmt.Println(s3) //[100 1 2 3 4 5]
	s3 = s3[:3]     //Slice后再Slice
	fmt.Println(s3) //[100 1 2]
	s3 = s3[2:]
	fmt.Println(s3) //[2]
	//Slice扩展
	arr3 := [...]int{0, 1, 2, 3, 4, 5, 6, 7}
	a1 := arr3[2:6]
	a2 := a1[3:5]
	fmt.Println(a1, a2) //[2 3 4 5] [5 6]  a1一共就4个元素  第5个元素取不到，6是怎么来的？
	//Slice可以向后扩展，不可以向前扩展
	//Slice里有三个变量
	//ptr：指向Slice开头的元素
	//len：Slice的长度
	//cap：整个arr，从ptr开始到结束
	// 只要不超过cap都能扩展
	// s[i]不可以超越len(s),向后扩展不可以超越底层数组cap(s)
	// 可以取得len(s)和cap(s)的值
	fmt.Printf("%v,%v,%v", a1, len(a1), cap(a1)) //[2 3 4 5],4,6
}
func updateSlice(s []int) {
	s[0] = 100
}
func printArray1(arr []int) {
	arr[0] = 100
	for i, v := range arr {
		fmt.Println(i, v)
	}
}
