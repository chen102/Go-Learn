package main

import "fmt"

func main() {
	arr := [...]int{0, 1, 2, 3, 4, 5, 6, 7}
	s1 := arr[2:6]
	fmt.Println(s1)
	s2 := s1[3:5]
	//向Slice添加元素
	fmt.Println(s2)
	s3 := append(s2, 1)
	fmt.Println(s3)
	fmt.Println(arr) //[0 1 2 3 4 5 6 1]	append将7变成了1
	//添加元素时如果超过cap，系统会重新分配更大的底层数组
	s4 := append(s3, 2)
	fmt.Println(s4)
	fmt.Println(arr) //[0 1 2 3 4 5 6 1] append的2去哪了？
	//在一个新的view里即s4不在是1对arr的一个view
	//由于值传递的关系，必须接受append的返回值
	//append一个元素，slice的变量都会改变 即必须用新的Slice来接这个append的返回值
	//s = append(s,val)

	//创建Slice
	var s []int //Zero value for slice is nil
	for i := 0; i < 10; i++ {
		printSlice(s) //cap不够了，就翻一倍
		// s=[] len=0 cap=0
		// s=[0] len=1 cap=1
		// s=[0 1] len=2 cap=2
		// s=[0 1 2] len=3 cap=4
		// s=[0 1 2 3] len=4 cap=4
		// s=[0 1 2 3 4] len=5 cap=8
		// s=[0 1 2 3 4 5] len=6 cap=8
		// s=[0 1 2 3 4 5 6] len=7 cap=8
		// s=[0 1 2 3 4 5 6 7] len=8 cap=8
		// s=[0 1 2 3 4 5 6 7 8] len=9 cap=16
		s = append(s, i)
	}
	fmt.Println(s)
	a1 := []int{0, 1, 2, 3, 4}
	fmt.Println(a1, len(a1), cap(a1)) //[0 1 2 3 4] 5 5
	a2 := append(a1, 5)
	fmt.Println(a2, len(a2), cap(a2)) //[0 1 2 3 4 5] 6 10

	//知道多大的Slice，值不知道
	a3 := make([]int, 16)     //长度是16
	a4 := make([]int, 10, 32) //还可以跟上cap
	printSlice(a3)
	printSlice(a4)
	// s=[0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0] len=16 cap=16
	// s=[0 0 0 0 0 0 0 0 0 0] len=10 cap=32

	//copying slice
	copy(a4, a2)    //将a2拷进a4
	fmt.Println(a4) //[0 1 2 3 4 5 0 0 0 0]

	//deleting elements from slice
	//没有删除的关键字，但可以用reslice
	//删掉[0 1 2 3 4 5 0 0 0 0]中的3
	//利用reslice和append删除中间的元素
	a4 = append(a4[:3], a4[4:]...) //a4[4:]... :第4个元素后面所有的元素
	printSlice(a4)                 //s= len=9 cap=32

	//一般删除元素用于去头去尾
	//删掉[0 1 2 4 5 0 0 0 0]的头尾
	a4 = a4[1:]
	a4 = a4[:len(a4)-1]
	printSlice(a4) //s=[1 2 4 5 0 0 0] len=7 cap=31
}
func printSlice(s []int) {
	fmt.Printf("s=%v len=%v cap=%v\n", s, len(s), cap(s))
}
