package main

import (
	"fmt"
	"net"
	"os"
)

func main() {
	conn, err := net.Dial("tcp", "192.168.110.221:8001")
	defer conn.Close()
	if err != nil {
		fmt.Println("conn:", err)
		return
	}
	addr := conn.LocalAddr().String()
	fmt.Println(addr, "conncet ok")
	go Write(conn)
	Read(conn)
}
func Write(conn net.Conn) {
	buf := make([]byte, 2048)
	for {
		n, err := os.Stdin.Read(buf)
		if err != nil {
			fmt.Println("Write:", err)
			return
		}
		conn.Write(buf[:n-1])
	}
}
func Read(conn net.Conn) {
	buf1 := make([]byte, 2048)
	for {
		n, err := conn.Read(buf1)
		if err != nil {
			fmt.Println("Read:", err)
			return
		}
		fmt.Println((string(buf1[:n])))
	}
	fmt.Println("read close")
	conn.Close()

}
