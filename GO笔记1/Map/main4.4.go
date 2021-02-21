package main

import "fmt"

func main() {
	//定义Map
	m := map[string]string{
		"name":   "chenxi",
		"course": "glang",
		"site":   "imooc",
	}
	m2 := make(map[string]int) //m2 == empty map
	var m3 map[string]int      //m3 == nil
	//在运算时empty和nil可以混用
	fmt.Println(m, m2, m3)
	fmt.Println(len(m)) //通过len获取元素的个数

	//Map的遍历
	//不保证遍历顺序
	for i, v := range m {
		fmt.Println(i, v) //这些key在Map里是无序的
	}
	//Map的操作
	//获取元素：m[key]
	name := m["name"]
	fmt.Println(name)
	//当取Map不存在的key时，Map依旧可以取值-zero value,获得value类型的初始值
	site := m["fffff"]
	fmt.Println(site)
	fmt.Println("---------")

	//怎么判断有没有key呢
	name1, ok := m["name"]
	fmt.Println(name1, ok) //chenxi true
	site1, ok := m["fffff"]
	fmt.Println(site1, ok) //false

	if name2, ojbk := m["name"]; ojbk {
		fmt.Println(name2)
	} else {
		fmt.Println("key does not exist")
	}
	if site2, ojbk := m["fffff"]; ojbk {
		fmt.Println(site2)
	} else {
		fmt.Println("key does not exist")
	}

	//删除元素
	delete(m, "course")
	course, ok := m["course"]
	fmt.Println(course, ok)
}

//map使用哈希表，必须可以比较相等
//除了slice,map,function的内建类型都可以作为key
//struct类型不包含上述字段，也可以作为key
