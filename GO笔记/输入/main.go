package main

import "fmt"

func main() {
	var a int
	fmt.Printf("请输入： ")
	// fmt.Scanf("%d", &a)		//阻塞等待用户输入
	fmt.Scan(&a) //更简便，自动匹配格式
	fmt.Println(a)
}
