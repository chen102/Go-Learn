//在循环中有两个关键操作break和continue，break操作是跳出当前循环，continue是跳过本次循环
package main

import (
	"fmt"
	"time"
)

func main() {
	for i := 1; i < 10; i++ {
		if i == 6 {
			// break	//1,2,3,4,5
			continue //1,2,3,4,5,7,8,9
		}
		fmt.Println(i)
	}
	c := 0
	for {
		c++
		time.Sleep(time.Second) //延时一秒
		if c%2 == 0 {           //跳过偶数
			continue
		} else if c == 11 { //到10退出循环
			break
		}
		fmt.Println(c)
	}
}

//break可用于for,switch,select,而continue仅能用于for循环
