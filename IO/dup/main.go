package main

import (
	"bufio"
	"fmt"
	"os"
)

//判断多个文件的重复行数
func main() {
	count := make(map[string]int)
	file := os.Args[1:]
	if len(file) == 0 {
		fmt.Println("请加文件名")
	}
	for _, age := range file {
		f, err := os.Open(age)
		if err != nil {
			fmt.Println(err)
			continue
		}
		countlenes(f, count)
		f.Close()
	}
	for line, n := range count {
		if n > 1 {
			fmt.Println(n, line)
		}

	}

}
func countlenes(f *os.File, count map[string]int) {
	input := bufio.NewScanner(f)
	for input.Scan() {
		count[input.Text()]++
	}
}
