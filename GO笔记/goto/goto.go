// goto可以用在任何地方，但不能夸函数使用
package main

import "fmt"

func main() {
	fmt.Println(1)
	goto End //End是用户起的名字，叫标签
	fmt.Println(2)
End:
	fmt.Println(3)
}
