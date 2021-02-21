package main

import "fmt"
import "sync"

//sync.Once
func SyncPackAge() {
	var once sync.Once
	for i := 0; i < 10; i++ {
		once.Do(func() {
			fmt.Println(i)
		}) //注意不用调用匿名函数了() 只执行了一次
	}
}

//sync.Map
func SyncPackAge1() {
	var users sync.Map
	users.Store(1, "hehe")
	users.Store(2, "aaa")
	value, _ := users.Load(1)
	fmt.Println("1:", value)
	users.Delete(1)
	_, ok := users.Load(1)
	if !ok {
		fmt.Println("空的")
	}
}

//sync.Pool 对象池
func SyncPackAge2() {
	pool := sync.Pool{
		New: func() interface{} {
			fmt.Println("New")
			return 1
		},
	}
	x := pool.Get()
	fmt.Println(x)
	pool.Put(x)
	x = pool.Get()
	x = pool.Get()

}
func SyncPackAge3() {
	var wg sync.WaitGroup //计数信号量
	for i := 0; i < 10; i++ {
		wg.Add(1)         //添加一个信号
		go Test_3(i, &wg) //传的是指针
	}
	wg.Wait() //等待信号量归零.否则阻塞
}
func Test_3(i int, wg *sync.WaitGroup) {
	fmt.Println(i)
	wg.Done() //处理一个信号
}
