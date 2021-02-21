package main

import "fmt"

type Test struct { //若想被别的文件引用，开头必须大写
	x, y int
}

func zhizheng() {
	i := 10
	fmt.Println(i)
	fmt.Println(&i)
	p := &i
	fmt.Println(*p)
	*p = 22
	fmt.Println(i)
}

type Vertex struct {
	x int
	y int
}

func test() {
	fmt.Println(Vertex{1, 2})
	v := Vertex{1, 2}
	fmt.Println(v)
	v.x = 44
	fmt.Println(v.x)
	p := &v
	p.y = 333
	fmt.Println(v)
	p = &Vertex{1, 2}
	fmt.Println(p)
}
func test1() {
	var a [5]int
	for i := 0; i < len(a); i++ {
		a[i] = i
	}
	fmt.Println(a)
}
func test2() {
	p := new(int)
	*p = 666
	fmt.Println(*p)
}
func test3() {
	//var a [10]int = [10]int{1, 2, 3, 4}
	arr := [10]int{1, 2, 3, 4}
	fmt.Println(arr)
	arr1 := [5]int{2: 55, 4: 222} //指定元素初始化
	fmt.Println(arr1)
}
func test4() {
	a := []int{1, 2, 3, 4, 5}
	s := a[0:3:5] //[low:high:max] len=high-low cap=max-low
	fmt.Println("s=", s)
	fmt.Println("len(s)=", len(s))
	fmt.Println("cap(s)=", cap(s))
	s = a[1:4:5]
	fmt.Println("s=", s)
	fmt.Println("len(s)=", len(s))
	fmt.Println("cap(s)=", cap(s))
}
func test5() {
	//[]里写了数字就是数组
	a := [5]int{} //数组 len和cap不变
	c := []int{}  //切片 []里面为空或...
	c = append(c, 11)
	fmt.Println(a)
	fmt.Println(c)
	slice1 := []int{1, 2, 3, 4}
	fmt.Println(slice1)
	//make(切片类型，长度，容量)
	slice2 := make([]int, 5, 10) //不指定容量的话，将长度设为容量
	fmt.Println(slice2)

}
func test6() {
	arr2 := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	s1 := arr2[:] //[0:len(arr2):len(arr2)] 不指定容量，和长度一样
	fmt.Println("s1=", s1)
	fmt.Println("len=", len(s1))
	fmt.Println("cap=", cap(s1))
	s2 := arr2[1:3:6] //[low:high:max] len=high-high cap=max-low
	fmt.Println("s2=", s2)
	fmt.Println("len=", len(s2))
	fmt.Println("cap=", cap(s2))
	s3 := s2[:5] //从0开始去5个元素
	fmt.Println(s3)
	s3[0] = 999
	fmt.Println(s3)
	fmt.Println(arr2)
	s4 := append(s3, 999)
	fmt.Println("s4=", s4)
	test6_1(s4) //引用传递
	fmt.Println("s4=", s4)

}
func test6_1(s []int) {
	s[0] = 555
}
func test7() {
	//var m1 map[int]string

	m2 := make(map[int]string, 3) //指定长度
	m2[1] = "hehe"
	m2[2] = "cdcd"
	m2[3] = "gggg" //长度自动扩容
	fmt.Println("len=", len(m2))
	m3 := map[int]string{1: "cc", 2: "rrr"}
	fmt.Println(m3)
	m4 := make(map[int][]string) //用切片作为map的值
	m4[1] = []string{"a", "b", "c"}
	m4[2] = append(m4[1], "d")
	fmt.Println(m4)
	for key, value := range m4 {
		fmt.Println(key, value)
	}
	delete(m4, 1) //删除key为1的内容
	fmt.Println(m4)
	test7_1(m4) //引用传递
	fmt.Println(m4)
}
func test7_1(m map[int][]string) {
	delete(m, 2)
}
func test8(s stu) {
	s.id = 666

}
func test9(s *stu) {
	s.id = 666
}
