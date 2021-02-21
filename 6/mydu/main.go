//版本1:使用range遍历channel 累加size
//版本2:增加了可选参数v：定时打印更新流 用了select监听tick，所以显式地判断fileSize是否关闭(没有使用range了)
//版本3:对每个目录创建一个新协程处理
//版本4:增加了临时退出
package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"sync"
	"time"
)

//工具函数,这个函数被调用时候会轮询退出状态
var done = make(chan struct{})

func cancelled() bool {
	select {
	case <-done:
		return true
	default:
		return false
	}
}
func main() {
	var verbose = flag.Bool("v", false, "show verbose progress message") //-v 参数返回一个bool指针(false)
	flag.Parse()                                                         //解析命令行参数
	roots := flag.Args()
	fmt.Println(roots)
	if len(roots) == 0 {
		roots = []string{"."} //若没有目录，以当前目录
	}
	//版本4
	go func() { //输入任意键关闭done
		os.Stdin.Read(make([]byte, 1))
		close(done) //广播机制：不往里面写，直接关闭会不断写0值，再利用轮询函数退出其他协程
	}()
	//版本3
	fileSize := make(chan int64)
	var n sync.WaitGroup //任务队列
	for _, root := range roots {
		n.Add(1)
		go walkDir1(root, &n, fileSize) //协程处理目录
	}
	go func() { //WaitGroup计数器为0关闭channel
		n.Wait()
		close(fileSize)
	}()
	//版本1
	//fileSize := make(chan int64)
	//go func() {
	//for _, root := range roots { //遍历命令行参数中的所有目录
	//walkDir(root, fileSize)
	//}
	//close(fileSize)
	//}()
	var nfiles, nbytes int64
	//版本1
	//for size := range fileSize { 主协程遍历filesize即一旦channel中有数据就读出来否则堵塞
	//nfiles++
	//nbytes += size
	//}
	//printDiskUsage(nfiles, nbytes)

	//版本2
	var tick <-chan time.Time
	if *verbose {
		tick = time.Tick(500 * time.Millisecond) //每500毫秒向tick写入
		fmt.Println("加参数了")
	}
loop: //标签break
	for {
		select {
		case <-done: //退出
			for range fileSize { //将fileSize中排空 避免调用walkDir后被阻塞
			}
		case size, ok := <-fileSize:
			if !ok {
				break loop //若fileSize中没有数据，返回循环出口
			}
			nfiles++
			nbytes += size
		case <-tick: //若加了参数v 每500毫秒生成打印事件
			printDiskUsage(nfiles, nbytes)
		}
	}
	printDiskUsage(nfiles, nbytes)
	fmt.Println("结束了")
}
func printDiskUsage(nfiles, nbytes int64) {
	fmt.Printf("%d files %.1f GB\n", nfiles, float64(nbytes)/1e9)
}
func walkDir(dir string, fileSize chan<- int64) { //只写channel
	for _, entry := range dirents(dir) { //返回目录下的文件列表并遍历
		if entry.IsDir() {
			subdir := filepath.Join(dir, entry.Name()) //若是目录，递归
			walkDir(subdir, fileSize)
		} else {
			fileSize <- entry.Size()
		}
	}
}
func walkDir1(dir string, n *sync.WaitGroup, fileSize chan<- int64) {
	defer n.Done()
	if cancelled() { //检查是否已退出
		return
	}
	for _, entry := range dirents2(dir) {
		if entry.IsDir() {
			n.Add(1)                                   //每扫描到一个目录，任务队列+1
			subdir := filepath.Join(dir, entry.Name()) //若是目录，递归
			go walkDir1(subdir, n, fileSize)           //创建新的协程处理目录
		} else {
			fileSize <- entry.Size()

		}
	}
}

func dirents(dir string) []os.FileInfo {
	entries, err := ioutil.ReadDir(dir) //返回目录信息的有序列表
	if err != nil {
		fmt.Fprintf(os.Stderr, "du1:%v\n", err)
		return nil
	}
	return entries
}

var sema = make(chan struct{}, 20)

func dirents1(dir string) []os.FileInfo {
	sema <- struct{}{} //缓冲，阻止他同时打开太多的文件
	defer func() {
		<-sema
	}()
	entries, err := ioutil.ReadDir(dir) //返回目录信息的有序列表
	if err != nil {
		fmt.Fprintf(os.Stderr, "du1:%v\n", err)
		return nil
	}
	return entries
}
func dirents2(dir string) []os.FileInfo {
	defer func() {
		<-sema
	}()
	select {
	case sema <- struct{}{}: //缓冲，阻止他同时打开太多的文件
	case <-done:
		return nil
	}
	entries, err := ioutil.ReadDir(dir) //返回目录信息的有序列表
	if err != nil {
		fmt.Fprintf(os.Stderr, "du1:%v\n", err)
		return nil
	}
	return entries
}
