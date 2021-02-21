package main

import "fmt"

func test01() {
	s1 := Student{Person{"c", 'm', 22}, 2, "iii"} //顺序初始化
	fmt.Println(s1)
	s2 := Student{Person: Person{name: "ff", age: 55}} //指定初始化
	fmt.Println(s2)
	s2.name = "www" //默认修改的是Student的name
	fmt.Println(s2)
	//显式调用
	s2.Person.name = "ccc"
	fmt.Println(s2)
	//非结构体函数
	s3 := Stu{Person{"c", 'm', 33}, 2, "tt"}
	fmt.Println(s3)
	s3.str = "tttt"
	s3.int = 55
	fmt.Println(s3)
	s3.Person = Person{"aghej", 'M', 666}
	fmt.Println(s3)
	//结构体指针类型匿名字段
	//直接初始化赋值
	s4 := Student1{&Person{"c", 'm', 22}, 2, "iii"}
	fmt.Println(s4.name, s4.sex, s4.age, s4.id)
	//先分配空间再赋值
	var s5 Student1
	s5.Person = new(Person)
	s5.name = "oiadbhg"
	fmt.Println(s5.name)

}
func test02() {
	var a long = 2
	s := a.Add(3) //变量名.函数(所需参数)
	fmt.Println(s)
}
func test03() {
	p := Person{"mike", 'n', 19}
	p.PrintInfo()
}
func test04() {
	p2 := Person{"sss", 'g', 44}
	p2.SetInfoValue("ttasdgt", 'f', 55) //值传递
	p2.PrintInfo()
	(&p2).SetInfoPointer("tttt", 'f', 22) //引用传递
	p2.PrintInfo()
}
func test05() {
	var a long = 5
	var c char = 'c'
	a.Print()
	c.Print()
}
func test07() {
	//指针变量的方法集
	p := &Person{"sgas", 'g', 54}
	p.SetInfoValue("cc", 'f', 44)      //(*p).SetInfoValue() 内部转换，把指针p，转成*p再调用
	(*p).SetInfoPointer("cc", 'g', 55) //p.SetInfoPointer() 内部转换..
	//不同变量的方法集
	p1 := Person{"ghg", 'f', 5}
	p1.SetInfoPointer("ff", 'g', 43) //(&p).SetInfoPointer()内部先把p转换为&p再调用

}
func test08() {
	s := Student{Person{"c", 'm', 22}, 2, "iii"}
	s.Print_Info()        //继承Person的方法(以被重写)
	s.Person.Print_Info() //显式调用继承方法
	pFunc := s.Print_Info //方法值 隐藏了接受者
	p := Person{"gg", 'g', 55}
	pFunc()
	//方法表达式
	f := (*Person).Print_Info //以类型操作，不用接受者
	f(&p)
	f1 := (Person).Print_Info
	f1(p)
}
func test09() {
	//定义接口变量
	var i my_interface
	s := stur1{1, 5}
	i = s //根据变量类型选择对应的方法
	i.test()
	s1 := stur2{2, 4}
	i = s1
	i.test()
	ohh(s)
	ohh(s1)
	//定义一个切片
	slic := make([]my_interface, 3)
	slic[0] = s
	slic[1] = s1
	fmt.Println(slic)
	var ii my_interface_big
	ii = s
	ii.test() //继承过来的方法
	ii.test1(66)
	i = ii //接口转换
	i.test()
}
func test10() {
	i := make([]interface{}, 3)
	i[0] = 555
	i[1] = "ggg"
	i[3] = 'c'
	for _, data := range i {
		if ok := data.(int); ok == true {
			fmt.Println("int")
		} else if ok := data.(string); ok == true {
			fmt.Println("string")
		} else if ok := data.(byte); ok == true {
			fmt.Println("byte")
		}

	}

}
