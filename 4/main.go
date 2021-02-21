package main

import "fmt"

//继承(代码复用)--匿名组合
type Person struct {
	name string
	sex  byte
	age  int
}
type Student struct {
	Person //匿名字段,Student包含了Person的所有字段
	id     int
	name   string //同名字段和Person同名了
}

func (p Student) Print_Info() {
	fmt.Println("以重写")
}

type str string

//非结构体函数
type Stu struct {
	Person
	int
	str
}

//结构体指针类匿名字段
type Student1 struct {
	*Person
	id   int
	name string
}

//方法-封装
type long int

func (tmp long) Add(other long) long {
	return tmp + other
}

//结构体类型添加方法
func (tmp Person) PrintInfo() {
	fmt.Println("tmp=", tmp)
}
func (p *Person) SetInfoPointer(n string, s byte, a int) {
	p.name = n
	p.sex = s
	p.age = a
}
func (p Person) SetInfoValue(n string, s byte, a int) {
	p.name = n
	p.sex = s
	p.age = a
}

type char byte

//只要接受类型不同，就是不同方法(golang不支持重载)
func (t char) Print() {
	fmt.Println(t)
}
func (u long) Print() {
	fmt.Println(u)
}
func (p Person) Print_Info() {
	fmt.Println(p.name, p.sex, p.age)
}

//定义接口类型
type my_interface interface {
	test()
}
type my_interface_big interface {
	my_interface
	test1(i int)
}
type stur1 struct {
	x, y int
}

func (s stur1) test() {
	fmt.Println("stur1")
}
func (s stur1) test1(i int) {
	fmt.Println(i)
}

type stur2 struct {
	x, y int
}

func (s stur2) test() {
	fmt.Println("stur2")
}

type stur3 struct {
	x, y int
}

func (s stur3) test() {
	fmt.Println("stur3")
}

//多态
func ohh(i my_interface) {
	i.test()
}
func oo(arg ...interface{}) { //接受任意类型多个参数

}
func main() {
	fmt.Println("xx")
	test01()
	test02()
	test03()
	test04()
	test05()
	test07()
	test08()
	test08()
	test09()
	var i interface{} //空接口万能类型，保存任意类型的值Println(a ...interface{})
	i = 1
	fmt.Println(i)
	//test10() 有问题
}
