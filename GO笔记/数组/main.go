package main

import "fmt"

func main() {
	arr := [...]int{3, 5, 6, 7, 1, 3}
	fmt.Println(arr[:])
	for _, v := range arr {
		fmt.Println(v)
	}
	var arr1 [10]int //[]里必须是常量，指定数组元素个数
	//赋值每个元素
	for i := 0; i < len(arr1); i++ {
		arr1[i] = i + 1
	}
	fmt.Println(arr1)

	//二维数组
	//有3个元素，每个元素又是一维数组[4]int
	b := [3][4]int{{1, 3, 4, 5}, {3, 45, 5, 3}, {3, 31, 32, 2}}
	fmt.Println(b)
	// 部分初始化
	r := [3][4]int{1: {1, 2, 3, 4}}
	fmt.Println(r)

	//数组支持比较，只支持==或！=，比较是不是每一个元素都一样，2个数组比较，数组类型要一样
	o := [3]int{3, 3, 2}
	p := [3]int{3, 3, 2}
	fmt.Println(o == p)
	var u [3]int
	u = o
	fmt.Println(u)

	//数组作函数参数，它是值传递
	//实参数组的每个元素给形参数组拷贝一份
	//形参的数组是实参数组的复制品
	g := [3]int{3, 4, 2}
	test(g)
	fmt.Println(g) //[3 4 2]

	//数组指针做函数参数
	test1(&g)
	fmt.Println(g) //[10 4 2]

}
func test(a [3]int) {
	a[0] = 10
	fmt.Println(a) //[10 4 2]
}
func test1(a *[3]int) {
	(*a)[0] = 10
	fmt.Println(*a) //[10 4 2]
}
