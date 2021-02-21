package main

import "fmt"
import "strconv"
import "strings"

func add(pk int, users map[string]map[string]string) {
	var (
		id   string = strconv.Itoa(pk)
		name string
		age  string
		tel  string
		addr string
	)
	fmt.Print("Name:")
	fmt.Scan(&name)
	fmt.Print("Age:")
	fmt.Scan(&age)
	fmt.Print("Tel:")
	fmt.Scan(&tel)
	fmt.Print("Addr:")
	fmt.Scan(&addr)
	users[id] = map[string]string{
		"id":   id,
		"name": name,
		"age":  age,
		"tel":  tel,
		"addr": addr,
	}
}
func Query(users map[string]map[string]string) {
	var q string
	fmt.Print("查询信息:")
	fmt.Scan(&q)
	for _, user := range users {
		if q == "" || strings.Contains(user["name"], q) || strings.Contains(user["age"], q) || strings.Contains(user["tel"], q) || strings.Contains(user["addr"], q) {
			fmt.Printf("%5s|%10s|%5s|%20s|%50s", user["id"], user["name"], user["age"], user["tel"], user["addr"])
			fmt.Println()
		}
	}
}
func Modfy(users map[string]map[string]string) {

	var (
		id   string
		name string
		age  string
		tel  string
		addr string
	)
	fmt.Print("修改id:")
	fmt.Scan(&id)
	user := users[id]
	if user == nil {
		fmt.Printf("没有用户")
		return
	}
	fmt.Printf("%5s|%10s|%5s|%20s|%50s\n", user["id"], user["name"], user["age"], user["tel"], user["addr"])
	fmt.Printf("你确定要修改(Y/N):")
	var q string
	fmt.Scan(&q)
	if q == "Y" || q == "y" {
		fmt.Print("Name:")
		fmt.Scan(&name)
		fmt.Print("Age:")
		fmt.Scan(&age)
		fmt.Print("Tel:")
		fmt.Scan(&tel)
		fmt.Print("Addr:")
		fmt.Scan(&addr)
		user["name"] = name
		user["age"] = age
		user["tel"] = tel
		user["addr"] = addr

		fmt.Print("修改成功")
		fmt.Printf("%5s|%10s|%5s|%20s|%50s\n", user["id"], user["name"], user["age"], user["tel"], user["addr"])
	} else if q == "N" || q == "n" {
		return
	} else {
		fmt.Printf("请输入Y/S")
	}

}
func Delete(users map[string]map[string]string) {
	var id string
	fmt.Printf("删除id:")
	fmt.Scan(&id)
	delete(users, "id")
}
func main() {
	users := make(map[string]map[string]string)
	id := 0
	P := "123abc"
	var s string
	fmt.Println("START")
	for i := 0; i < 3; i++ {
		fmt.Print("请输入密码:")
		fmt.Scan(&s)
		if s == P {
			for {
				var op int
				fmt.Print(`
1.Add User
2.Modfy User
3.Delete User
4.Query User
5.Quit
		`)
				fmt.Scan(&op)
				if op == 1 {
					id++
					add(id, users)
				} else if op == 2 {
					Modfy(users)
				} else if op == 3 {
					Delete(users)
				} else if op == 4 {
					Query(users)
				} else if op == 5 {
					break
				} else {
					fmt.Println("ERROR1:请重新输入")
				}
			}

		}

	}
	fmt.Println("密码错误多次，程序退出")
}
