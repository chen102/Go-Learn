package main

import (
	"bufio"
	"encoding/csv"
	"encoding/gob"
	"fmt"
	"io"
	"os"
	"strings"
)

//标准输入输出
func test7() {
	fmt.Println("xxx")
	os.Stdout.Write([]byte("xxx"))
	bytes := make([]byte, 5)
	n, err := os.Stdin.Read(bytes)
	fmt.Println(n, err, bytes)
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	fmt.Println(scanner.Text())
	fmt.Fprintf(os.Stdout, "%T", 1)
}

type User struct {
	ID       int
	Name     string
	Birthday string
	Tel      string
	Addr     string
}

//gob
//把内存中的数据持久到文件中
func test8() {
	users := map[int]User{
		1: User{1, "a", "1", "1", "1"},
		2: User{2, "b", "1", "1", "1"},
		3: User{2, "c", "1", "1", "1"},
	}
	file, err := os.Create("ohhh.gob")
	if err == nil {
		defer file.Close()
		//持久化
		encoder := gob.NewEncoder(file)
		encoder.Encode(users)
		fmt.Println(users)
	}
}

func test9() {
	users := map[int]User{}
	file, err := os.Open("ohhh.gob")
	if err == nil {
		defer file.Close()
		//反持久化
		decoder := gob.NewDecoder(file)
		decoder.Decode(&users)
		fmt.Println(users)
	}
}

//第二次持久化的方式csv
//csv编码
func test10() {
	file, err := os.Create("tt.csv")
	defer file.Close()
	if err != nil {
		fmt.Println(err)
	} else {
		writer := csv.NewWriter(file)
		writer.Write([]string{"编号", "名字", "性别"})
		writer.Write([]string{"1", "a", "F"})
		writer.Write([]string{"2", "b", "N"})
		writer.Flush()

	}
}

//csv解码
func test11() {
	file, err := os.Open("tt.csv")
	defer file.Close()
	if err != nil {
		fmt.Println(err)
	} else {
		reader := csv.NewReader(file)
		for {
			line, err := reader.Read()
			if err != nil {
				if err != io.EOF {
					fmt.Println(err)
				}
				return
			} else {
				fmt.Println(line)
			}

		}

	}

}

//处理流数据，直接在内存中处理
func test12() {
	reader := strings.NewReader("abcdef123456")
	bytes := make([]byte, 3)
	for {
		n, err := reader.Read(bytes)
		if err != nil {
			if err != io.EOF {
				fmt.Println(err)
			}
			break
		} else {
			fmt.Println(bytes[:n])
		}
	}
}

//判断文件重复的行
func test13() {
	counts := make(map[string]int)
	file, err := os.Open("test3.txt")
	defer file.Close()
	if err != nil {
		fmt.Println(err)
		return
	} else {
		input := bufio.NewScanner(file)
		for input.Scan() {
			counts[input.Text()]++

		}
	}
	for line, n := range counts {
		if n > 1 {
			fmt.Println(n, line)
		}
	}
}
