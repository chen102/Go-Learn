package main

import "fmt"

import "bufio"
import "net"
import "os"

//import "net/http"
//import _ "net/http/pprof"

func main() {
	//if err := http.ListenAndServe("127.0.0.1:8002", nil); err != nil {
	//fmt.Printf("start pprof failed")
	//}
	listener, err := net.Listen("tcp", "192.168.110.221:8001")
	if err != nil {
		fmt.Println("listen err:", err)
		os.Exit(-1)
	}
	for {
		conn, ok := AcceptRequest(listener)
		if ok {
			go HandleConn(conn)
		} else {
			continue
		}

	}
}
func AcceptRequest(listener net.Listener) (net.Conn, bool) {
	conn, err := listener.Accept()
	if err != nil {
		fmt.Println("conn err:", err)
	}
	return conn, true
}
func HandleConn(conn net.Conn) {
	clientAddr := conn.RemoteAddr()
	fmt.Println(clientAddr, "已连接")
	defer conn.Close()
	go func() {
		reader := bufio.NewReader(conn)
		for {
			n, err := reader.ReadString('\n')
			if err != nil {
				fmt.Println("Read err:", err)
				break
			}
			if n == "quit\n" {
				fmt.Println(clientAddr, "已断开")
				break
			}
			fmt.Printf("[%s]:%s", clientAddr, n)
		}
	}()
	writer := bufio.NewWriter(conn)
	input := bufio.NewReader(os.Stdin)
	for {
		fmt.Printf("我：")
		s, _ := input.ReadString('\n')
		_, err := writer.WriteString(s)
		if err != nil {
			fmt.Println("Write err", err)
			return
		}
		writer.Flush()
	}
}
