package main

import "fmt"

func main() {
	//定义一维数组
	var arr1 [5]int
	arr2 := [3]int{1, 3, 5}
	arr3 := [...]int{3, 4, 5, 6, 7, 4} //[...]让编译器数有几个int
	fmt.Println(arr1, arr2, arr3)
	//定义二维数组
	var grid [4][5]int //4个长度为5的int
	fmt.Println(grid)
	//数组的遍历
	for i := 0; i < len(arr3); i++ {
		fmt.Println(arr3[i])
	}
	//range关键字,同时获得数组的元素及元素的第几个值
	for i := range arr3 {
		fmt.Println(arr3[i])
	}
	for i, v := range arr3 {
		fmt.Println(i, v)
	}

	//只想要值
	for _, v := range arr3 {
		fmt.Println(v)
	}

	//数组的值的和
	sum := 0
	for _, v := range arr3 {
		sum += v
	}
	fmt.Println(sum)

	//数组的最后一个元素及值
	maxi := -1
	maxValue := -1
	for i, v := range arr3 {
		if v > maxValue {
			maxi, maxValue = i, v
		}
	}
	fmt.Println(maxi, maxValue)

	printArray(arr3) //cannot use arr3 (type [6]int) as type [5]int in argument to printArray
	printArray(arr2) //cannot use arr2 (type [3]int) as type [5]int in argument to printArray
	// 类型不同，不能传进去
	arr5 := [5]int{1, 2, 3, 4, 5}
	printArray(arr5)
	fmt.Println(arr5) //arr【0】并没有改变

	printArray1(&arr5)
	fmt.Println(arr5)  //arr[0]变了
	printArray1(&arr2) //cannot use &arr2 (type *[3]int) as type *[5]int in argument to printArray1
	//用指针，类型不同也传不进去
}

//数组是值类型
func printArray(arr [5]int) { //必须规定数组的长度
	arr[0] = 100
	for i, v := range arr {
		fmt.Println(i, v)
	}

}

//指向数组的指针
func printArray1(arr *[5]int) {
	arr[0] = 100
	for i, v := range arr {
		fmt.Println(i, v)
	}
}

//很麻烦，每次都要用指针，而且要知道元素的个数
//在go语言中一般不直接使用数组
//go语言使用切片
