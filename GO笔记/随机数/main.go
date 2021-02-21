package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	//设置种子，只需一次
	//如果种子参数一样，每次运行程序产生是随机数都一样
	// rand.Seed(666)
	//拿当前时间作为参数
	rand.Seed(time.Now().UnixNano())
	fmt.Println(rand.Int())     //随机数很大
	fmt.Println(rand.Intn(100)) //限定在100以内

	//冒泡排序
	var arr [10]int
	n := len(arr)
	for i := 0; i < n; i++ {
		arr[i] = rand.Intn(100)
	}
	fmt.Println(arr)

	for i := 0; i < n-1; i++ {
		for j := 0; j < n-1; j++ {
			if arr[j] > arr[j+1] {
				arr[j], arr[j+1] = arr[j+1], arr[j]
			}
		}
	}
	fmt.Println(arr)
}
