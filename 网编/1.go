package main

import (
	"bufio"
	"fmt"
	"net"
	"strings"
)

func test1() {
	//监听
	listener, err := net.Listen("tcp", "192.168.110.221:8001")
	fmt.Printf("Type:%T", listener)
	if err != nil {
		fmt.Println(err)
		return
	}
	//阻塞等待用户链接
	conn, err := listener.Accept()
	if err != nil {
		fmt.Println(err)
		return
	}
	//接受用户的请求
	buf := make([]byte, 1024)
	n, err1 := conn.Read(buf)
	if err1 != nil {
		fmt.Println(err1)
		return
	}
	fmt.Println(string(buf[:n]))
}
func test2() {
	listener, err := net.Listen("tcp", "192.168.110.221:8001")
	if err != nil {
		fmt.Println(err)
		return
	}
	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println(err)
			return
		}
		go HandleConn(conn)
	}
}
func HandleConn(conn net.Conn) {
	defer conn.Close()
	//获取客户端的网络地址信息
	addr := conn.RemoteAddr().String()
	fmt.Println(addr, "conncet ok")
	buf := make([]byte, 2024)
	conn.Write([]byte("Server:输入quit退出"))
	for {
		n, err := conn.Read(buf)
		if err != nil {
			fmt.Println(err)
			return
		}
		fmt.Printf("[%s]:%s\n", addr, string(buf[:n]))
		if string(buf[:n]) == "quit" {
			fmt.Println(addr, "已经退出")
			return
		}
		conn.Write([]byte(strings.ToUpper(string(buf[:n])))) //转换成大写
	}
}

//测试
func HandleConn1(conn net.Conn) {
	defer conn.Close()
	reader := bufio.NewReader(conn)
	writer := bufio.NewWriter(conn)
	line, err := reader.ReadString('\n')
	fmt.Println(line, err)
	writer.WriteString("哈哈" + "\n")
	writer.Flush() //缓冲io不要忘了flush
}
