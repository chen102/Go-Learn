package main

import "fmt"

type stu struct {
	id   int
	name string
}

func main() {
	fmt.Println("vim-go")
	zhizheng()
	test()
	test1()
	test2()
	test3()
	test4()
	test5()
	test6()
	test7()
	var s1 stu = stu{1, "abc"} //顺序初始化，每个成员必须初始化
	fmt.Println(s1)
	var s2 stu = stu{id: 1}
	fmt.Println(s2)
	var s3 *stu = &stu{1, "abc"}
	fmt.Println(*s3)
	var s4 stu
	s4.id = 5
	fmt.Println(s4.id)
	//指针方式调用结构体变量
	var s stu
	p := &s //需要先定义指针变量指向结构体变量
	p.id = 5
	fmt.Println(p.id)
	//或者使用new
	s5 := new(stu)
	s5.id = 5
	fmt.Println(s5.id)
	test8(s2) //值传递
	fmt.Println(s2.id)
	test9(&s2)
	fmt.Println(s2.id)

	test01()
}
