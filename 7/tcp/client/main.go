package main

import "fmt"
import "os"
import "net"
import "io"

func main() {
	list := os.Args
	if len(list) != 2 {
		fmt.Println("请在后面加上文件")
		return
	}
	fileName := list[1]
	info, err := os.Stat(fileName)
	if err != nil {
		fmt.Println("err:", err)
		return
	}
	conn, err1 := net.Dial("tcp", "192.168.110.221:8001")
	if err1 != nil {
		fmt.Println("conn err:", err1)
		return
	}

	_, err2 := conn.Write([]byte(info.Name()))
	if err2 != nil {
		fmt.Println("conn Write err:", err2)
		return
	}
	var n int
	buf := make([]byte, 1024)
	n, err3 := conn.Read(buf)
	if err3 != nil {
		fmt.Println("conn Read err:", err3)
		return
	}
	if "ok" == string(buf[:n]) {
		SendFile(fileName, conn)
	}
}
func SendFile(fileName string, conn net.Conn) {
	f, err := os.Open(fileName)
	if err != nil {
		fmt.Println("err:", err)
		return
	}
	defer f.Close()
	buf := make([]byte, 4*1024)
	for {
		n, err := f.Read(buf)
		if err != nil {
			if err == io.EOF {
				fmt.Println("发送成功")
			} else {
				fmt.Println("f Read err:", err)

			}
			return

		}
		conn.Write(buf[:n])
	}

}
