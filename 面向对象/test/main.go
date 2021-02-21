package main

import "fmt"

type Class struct {
	classid int
	tid     int
}
type Student struct {
	Class
	uid int
}

func (s Student) PrinfStudentInfo() {
	fmt.Println("学号为:", s.uid)
}
func (c Class) PrinfClassInfo() {
	fmt.Println("班级为", c.classid)
	fmt.Println("班主任为:", c.tid)
}

type SetInfoer interface {
	setStudentInfo()
}

func (c *Class) setStudentInfo() {
	c.classid++
	c.tid++
}
func (s *Student) setStudentInfo() {
	s.uid++
}
func SetINFO(i SetInfoer) {
	i.setStudentInfo()
}
func main() {
	fmt.Println("继承------------------------------------------------------------")
	s := Student{Class{1, 1}, 1}
	s.PrinfStudentInfo()
	s.PrinfClassInfo()
	fmt.Println("接口-------------------------------------------------------------")
	var i SetInfoer
	i = &s
	i.setStudentInfo()
	s.PrinfClassInfo()
	s.PrinfStudentInfo()
	i = &s.Class
	i.setStudentInfo()
	s.PrinfClassInfo()
	s.PrinfStudentInfo()

	fmt.Println("多态---------------------------------------------------------------")
	SetINFO(&s)
	s.PrinfClassInfo()
	s.PrinfStudentInfo()
	SetINFO(&s.Class)
	s.PrinfClassInfo()
	s.PrinfStudentInfo()
	fmt.Println("方法值和方法表达式-------------------------------------------------")
	PrintfClass := s.PrinfClassInfo
	PrintfClass()
	f := Student.PrinfStudentInfo
	f1 := (*Student).setStudentInfo
	f1(&s)
	f(s)

}
