package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

//1加到100
func xh() {
	sum := 0
	for i := 1; i <= 100; i++ { //for的条件里不需要括号
		sum += i
		fmt.Println(sum)
	}
}

//十进制转2进制  除2，对2取模，逆序
func convertToBin(n int) string {
	result := ""
	for ; n > 0; n /= 2 {
		lsb := n % 2
		result = strconv.Itoa(lsb) + result //strconv.Itoa:将lsb转成字符串   将对2取模的数字往前加--不用逆序了
	}
	return result
}

//读文件（一行一行的读）
func printFile(filename string) {
	file, err := os.Open(filename) //打开这个文件
	if err != nil {
		panic(err) //报错中断
	}
	scanner := bufio.NewScanner(file) //bufio.Scanner:扫描模块 一行一行的读
	for scanner.Scan() {              //只有结束条件
		fmt.Println(scanner.Text())
	}
}

//死循环
func forerer() { //什么都不加就是死循环  go语言经常要用到死循环
	for {
		fmt.Println("abc")
	}
}

func main() {
	xh()
	fmt.Println(convertToBin(66), convertToBin(88))
	printFile("abc.txt")
	forerer()
}
