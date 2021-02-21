package test //	不同目录，包名不一样

import "fmt"

func Test1() { //被main函数调用，包函数名字如果首字母是小写，无法让别人调用，要想别人能调用，必须首字母大写
	fmt.Println("www")
}
func Test2() {
	fmt.Println("hhh")
}
func init() { //只要这个test包 被导入就会执行
	fmt.Println("this is test init")
}
