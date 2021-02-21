//不同级目录
package main

import "test"

//	可以通过 import _ "test" 来调用init函数
func main() {
	test.Test1() //调用不同包里面的函数，格式：包名.函数名（）
	test.Test2()
}
