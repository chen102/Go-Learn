package main

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"
	"strconv"
	"time"
)

func test1() {
	path := "test.txt"
	file, err := os.Open(path) //返回一个文件指针，和错误
	if err != nil {
		fmt.Println(err)
	} else {
		bytes := make([]byte, 20) //读20个元素
		for {                     //循环，每次读取部分内容
			n, err := file.Read(bytes)
			if err == io.EOF {
				break
			} else if err != nil {
				fmt.Println(err)
			} else {
				fmt.Println(string(bytes[:n]))
			}

		}
		file.Close()
	}
}

//写覆盖调原来的
func test2() {
	path := "test2.txt"
	file, err := os.Create(path)
	if err != nil {
		fmt.Println(err)
		return
	} else {
		file.Write([]byte("adgbihjadoigh"))
		file.WriteString("315135")
		file.Close()
	}
}

//文件的追加OpenFile
func test3() {
	path := "log.txt"
	file, err := os.OpenFile(path, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0777) //os.ModePerm:0777权限
	//O_RDONLY int = syscall.O_RDONLY // 只读模式打开文件
	//O_WRONLY int = syscall.O_WRONLY // 只写模式打开文件
	//O_RDWR   int = syscall.O_RDWR   // 读写模式打开文件
	//O_APPEND int = syscall.O_APPEND // 写操作时将数据附加到文件尾部
	//O_CREATE int = syscall.O_CREAT  // 如果不存在将创建一个新文件
	//O_EXCL   int = syscall.O_EXCL   // 和O_CREATE配合使用，文件必须不存在
	//O_SYNC   int = syscall.O_SYNC   // 打开文件用于同步I/O
	//O_TRUNC  int = syscall.O_TRUNC  // 如果可能，打开时清空文件
	if err != nil {
		fmt.Println(err)
	} else {
		//_, err := file.WriteString("nackjdngjnglkdang\n")
		log.SetOutput(file)                                 //日志重定向
		log.SetPrefix("users:")                             //日志前缀
		log.SetFlags(log.Flags() | log.Lshortfile)          //文件路径也加上
		log.Print(strconv.FormatInt(time.Now().Unix(), 10)) //日志时间
		if err != nil {
			fmt.Println(err)
		}
		file.Close()
	}
}

//文件指针
//文件开始 0  os.SEEK_SET
//当前位置 1  os.SEEK_CUR
//文件末尾 2  os.SEEK_END
func test4() {
	file, _ := os.Open("test.txt")
	defer file.Close()
	s := make([]byte, 100)
	fmt.Println(file.Seek(0, 1)) //当前文件指针位置
	file.Seek(0, 2)              //将文件偏移量设置成文件末尾
	fmt.Println(file.Seek(0, 1)) //当前文件指针位置
	n, err := file.Read(s)
	fmt.Println(string(s[:n]))
	fmt.Println(err)
	file.Seek(0, 0) //偏移量，相对位置
	n, err = file.Read(s)
	fmt.Println(string(s[:n]))
	fmt.Println(err)
}

//一次性读写ioutil
func test5() {
	path := "test.txt"
	file, err := os.Open(path)
	defer file.Close()
	if err != nil {
		fmt.Println(err)
	} else {
		byte, err := ioutil.ReadAll(file)
		if err != nil {
			fmt.Println(err)
		} else {
			fmt.Println(string(byte[:]))
		}

	}

}

//更简单
func test5_1() {
	path := "test.txt"
	byte, err := ioutil.ReadFile(path)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(string(byte[:]))
}

//bufio
//scanner()
func test6() {
	path := "log.txt"
	file, err := os.Open(path)
	defer file.Close()
	if err == nil {
		scanner := bufio.NewScanner(file)
		for scanner.Scan() { //扫描，判断文件是否到末尾了
			fmt.Println(scanner.Text()) //拿取文件的一行内容
		}
	}
}

//Reader()
func test6_1() {
	path := "test.txt"
	file, err := os.Open(path)
	defer file.Close()
	if err != nil {
		fmt.Println(err)
		return
	}
	reader := bufio.NewReader(file)
	bytes := make([]byte, 5)
	for {
		_, err := reader.Read(bytes) //自动检查文件流的换行符\n进行截断
		if err != nil {
			if err != io.EOF {
				fmt.Println(err)
			}
			return
		}
		fmt.Println(string(bytes))

	}
}

//也可以指定进行截断 //ReadSlice()
func test6_2() {
	path := "test.txt"
	file, err := os.Open(path)
	defer file.Close()
	if err != nil {
		fmt.Println(err)
		return
	}
	reader := bufio.NewReader(file)
	for {
		line, err := reader.ReadSlice('\n')
		if err != nil {
			if err != io.EOF {
				fmt.Println(err)
			}
			return
		}
		fmt.Println(string(line))

	}
}
func test6_3() {
	path := "test.txt"
	file, err := os.Open(path)
	defer file.Close()
	if err != nil {
		fmt.Println(err)
		return
	}
	reader := bufio.NewReader(file)
	for {
		line, err := reader.ReadString('\n') //可直接读成string类型，下面不用转换了
		if err != nil {
			if err != io.EOF {
				fmt.Println(err)
			}
			return
		}
		fmt.Println(line)

	}
}
func test6_4() {
	path := "test00.txt"
	file, err := os.Create(path)
	defer file.Close()
	if err != nil {
		fmt.Println(err)
		return
	}
	writer := bufio.NewWriter(file)
	writer.WriteString("okok")
	writer.Flush() //刷新
}
