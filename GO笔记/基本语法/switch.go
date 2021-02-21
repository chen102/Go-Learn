package main

import "fmt"

func main() {
	// num := 1
	// fmt.Println("你想去几楼？")
	// fmt.Scan(&num)
	// switch num {
	// case 1:
	// 	fmt.Println("一楼")		//go语言不用break 自动跳出除非用fallthrough关键字
	// case 2:
	// 	fmt.Println("二楼")
	// case 3:
	// 	fmt.Println("三楼")
	// case 4:
	// 	fmt.Println("四楼")
	// default:
	// 	fmt.Println("没有")
	// }
	switch gr := 70; { //可以在switch中初始化
	case gr > 90:
		fmt.Printf("ooook")
	case gr > 70:
		fmt.Printf("oook")
	case gr > 50:
		fmt.Printf("ook")
	case gr > 0:
		fmt.Printf("ok")
	}
	fmt.Println(gr) //undefined: gr	gr没有定义 和if一样在语句中初始化，作用域只在语句中
}
