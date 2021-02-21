package main

//go语言的指针不能运算
import "fmt"

func zz() {
	var a int = 2
	var pa *int = &a //定义指针指向int 为a所在的地址
	*pa = 3          //a的值改为3
	fmt.Println(a)
}

//用函数交换两个值
func swap(a, b int) {
	b, a = a, b
}
func swap1(a, b *int) {
	*b, *a = *a, *b //b指向的内容等于a指向的内容，a指向的内容等于b指向的内容
}

//第二种方法
func swap2(a, b int) (int, int) {
	return b, a
}
func main() {
	zz()
	a, b := 3, 4
	swap(a, b) //交换不了因为是值传递
	fmt.Println(a, b)

	swap1(&a, &b)
	fmt.Println(a, b)

	a, b = swap2(a, b)
	fmt.Println(a, b)
}

//值传递和引用传递
//值传递：函数去main函数将值拷贝一份到自己函数内部进行修改，main函数的值不动
//引用传递：函数和main函数引用同一个值，函数的值变了，main的值也跟着变
//
//go语言使用值传递
//每次调用函数，函数都要拷贝，性能下降
//值传递和指针配合使用
//	函数用指针
//		func main(){
//		var a int
//		}
//		func f(pa *int){}
//		main函数的值的地址（&a）和函数的指针(pa)同时指向main的值
//			修改pa的内容，即修改了main的值
//				通过指针的传递来实现一种相当于引用传递的效果
