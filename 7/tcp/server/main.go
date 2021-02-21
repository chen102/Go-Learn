package main

import "fmt"
import "net"
import "os"
import "io"

func main() {
	listener, err := net.Listen("tcp", "192.168.110.221:8001")
	if err != nil {
		fmt.Println("Listen err:", err)
		return
	}
	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println(err)
			return

		}
		fileName, temp := recAccepte(conn)
		fmt.Println(fileName, temp)
		if temp {
			go func(fileName string) { //开始接收文件
				defer conn.Close()
				buf := make([]byte, 4*1024)
				file, err := os.Create(fileName)
				if err != nil {
					fmt.Println("Create file err:", err)
					return
				}
				defer file.Close()
				for {
					n, err := conn.Read(buf)
					if err != nil {
						if err == io.EOF {
							fmt.Println("接收成功")

						} else {
							fmt.Println("Read err", err)
						}
						return
					}
					file.Write(buf[:n]) //写入文件
				}

			}(fileName)
		} else {
			continue //等待下次传输文件请求
		}

	}
}
func recAccepte(conn net.Conn) (string, bool) { //处理传输请求
	var temp string
	addr := conn.RemoteAddr().String() //获取客户的地址
	buf := make([]byte, 1024)
	n, err := conn.Read(buf)
	if err != nil {
		fmt.Println("Read err", err)
	}
	fmt.Println(addr, string(buf[:n]))
	fmt.Printf("yes|else:")
	fmt.Scan(&temp)
	if temp == "yes" { //是否接收文件
		conn.Write([]byte("ok"))
		return string(buf[:n]), true
	} else {
		return "", false
	}

}
