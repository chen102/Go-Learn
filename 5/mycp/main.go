package main

import "fmt"
import "os"
import "io"

func main() {
	fmt.Println("vim-go")
	list := os.Args
	if len(list) != 3 {
		fmt.Println("请加上源文件和目标文件")
		return
	}
	srcFileName := list[1]
	dstFileName := list[2]
	if srcFileName == dstFileName {
		fmt.Println("源文件和目标文件名不能相同")
		return
	}
	b, err := os.Open(srcFileName)
	if err != nil {
		fmt.Println("err:", err)
		return
	}
	d, err1 := os.Create(dstFileName)
	if err1 != nil {
		fmt.Println("err:", err1)
		return
	}
	defer b.Close()
	defer d.Close()

	buf := make([]byte, 4*1024)
	for {
		n, err2 := b.Read(buf)
		if err2 != nil {
			fmt.Println("err2", err2)
			if err2 == io.EOF {
				break
			}
		}
		d.Write(buf[:n])
	}

}
