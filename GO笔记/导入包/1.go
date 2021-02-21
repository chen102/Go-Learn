package main

//导入包的两种方法
/*import "fmt"
import "os"

import (
	"bufio"
	"time"
)
*/

//.操作
/*import . "fmt" //调用函数，无需通过包名
func main() {
	Println("hahaha")	//不建议使用.操作  因为可能和自己设置的变量重名
}
*/

//给包名起别名
 /*import io "fmt"

 func main() {
 	io.Println("hahaha")
 }
*/

//忽略此包
import _ "fmt"
func main() {
	
	}
