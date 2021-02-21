package main

import "fmt"
import "net"
import "bufio"
import "os"

func main() {
	conn, err := net.Dial("tcp", "192.168.110.221:8001")
	if err != nil {
		fmt.Println("conn err:", err)
	}
	fmt.Println("输入quit退出")
	go Reader(conn)
	Writer(conn)
	conn.Close()

}
func Reader(conn net.Conn) {
	reader := bufio.NewReader(conn)
	Server_addr := conn.RemoteAddr().String
	for {
		n, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println("Read err:", err)
			return
		}
		fmt.Printf("[%s]:%s\n", Server_addr, n)
	}
}
func Writer(conn net.Conn) {
	writer := bufio.NewWriter(conn)
	input := bufio.NewReader(os.Stdin)
	for {
		fmt.Printf("我：")
		//input.Scan()
		//_, err := writer.WriteString(input.Text() + "\n")
		s, _ := input.ReadString('\n')
		_, err := writer.WriteString(s)
		if err != nil {
			fmt.Println("Write err", err)
			return
		}
		writer.Flush()
		if s == "quit\n" {
			return
		}
	}
}
