/*
bool,string   布尔和字符串
(u)int,(u)int8,(u)int16,(u)int32,(u)int64,uintptr   (u):无符号  不规定长度，跟着操作系统来   ptr：指针
byte，rune   rune：是go语言的字符型32位
float32，float64，complex64，complex128   complex：复数 complex64：32位 complex128：64位
*/

package main

import "fmt"
import "math"
import "math/cmplx" //为复数提供基本的常量和数学函数

func test() {
	c := 3 + 4i               //定义复数 不能用4*i，i会被认为是变量
	fmt.Println(cmplx.Abs(c)) //取绝对值（取模）
}

// 欧拉公式  e^iπ+1
func eulex() {
	//cmplx.Pow(math.E,i*math.Pi)+1）	这里的i会被认为是变量
	fmt.Println(
		cmplx.Pow(math.E, 1i*math.Pi)+1, //1i 编辑器就知道是虚数了
		cmplx.Exp(1i*math.Pi)+1)         //cmplx.Exp:e的多少次方
	//(0+1.2246467991473515e-16i)  为什么不是0  复数的实部和虚部分别是两个float

	fmt.Printf("%.3f\n", //只要要小数点后3位
		cmplx.Exp(1i*math.Pi)+1)
}

//强制类型转换 （go语言没有隐式类型转换）
func triangle() {
	var a, b int = 3, 4
	// var c int //Sqrt传回来是float64
	var c float64
	//或者直接将传回来的值强制转换成int
	// c = int(math.Sqrt(float64(a*a + b*b)))

	// c = math.Sqrt(a*a + b*b)    cannot use math.Sqrt(a * a + b * b) (type float64) as type int in assignment
	// Sqrt需要float64但传进去的是int，int不能隐式的转成float64   必须强制转换为float64
	c = math.Sqrt(float64(a*a + b*b))

	var d int = int(math.Sqrt(float64(a*a + b*b)))
	fmt.Println(c, d)
}

func main() {
	//test()
	//eulex()
	triangle()
}
