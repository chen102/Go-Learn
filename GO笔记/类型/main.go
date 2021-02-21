package main

import "fmt"

func main() {
	//布尔类型
	// var a bool
	// a = true
	// b := false
	// var c bool	//没有初始化，默认为false
	// fmt.Println(a, b, c)

	//浮点型
	// var f1 float32
	// f1 = 3.14
	// f2 := 3.14		//默认是float64
	// fmt.Println(f1, f2)
	// fmt.Printf("f2 type is %T", f2)
	// //float64存储小数比float32更准确

	//字符类型
	// a := 97
	// ch := 'a' //字符，单引号
	// fmt.Printf("%c %d\n", a, ch)
	// fmt.Printf("大写：%c 小写：%c ", 'A', 'a')
	// fmt.Printf("大写转小写：%c\n", 'A'+32)
	// fmt.Printf("小写转大写：%c\n", 'a'-32)

	//字符串
	// var str1 string
	// str1 = "abc"
	// str2 := "ccc"
	// fmt.Println(str1, str2)
	// fmt.Printf("str2 type is %T", str2)
	// //内建函数，len（）可以测字符串的长度，有多少个字符
	// fmt.Println("\nlen(str2)= ", len(str2))

	//字符
	//单引号，往往都只有一个字符，转义字符除外
	//字符串
	//双引号，字符串有一个或多个字符组成，字符串都是隐藏了一个结束符'\0'
	// str := "hello" //只想操作字符串的某个字符，从0开始操作
	// fmt.Printf("str[0]= %c  str[1]=%c \n", str[0], str[1])

	//复数类型
	// var t complex128
	// t = 3 + 4i
	// d := 3 + 4i //默认是complex128
	// fmt.Println(t)
	// fmt.Printf("d type is %T\n", d)
	// //通过内建函数real（），和imag取实部和虚部
	// fmt.Println(real(d), imag(d))
}
