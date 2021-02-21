package main

import "fmt"
import "sync"
import "time"
import "os"

var balance int

//竞争条件-数据竞争
//无论任何时候，只要有两个goroutine并发访问同一变量，且至少其中的一个是写操作的时候就会发生数据竞争

func Deposit(amout int) {
	balance += amout
}
func Balance() int {
	return balance
}
func bank() {
	s := make(chan int)
	go func() {
		Deposit(200)
		fmt.Println(Balance())
		s <- 1
	}()

	go Deposit(100)
	<-s
}

//3种方式避免数据竞争(避免并发访问共享变量)
//1.不要写变量 2.避免从多个goroutine访问变量 3.互斥锁

//2避免从多个goroutine访问变量
//只使用一个channel来发送指定的goroutine请求来查询更新
var deposits = make(chan int)
var balances = make(chan int)

func Deposit1(amout int) {
	deposits <- amout
}
func Balance1() int {
	return <-balances
}

//monitor goroutine监控协程:提供对一个指定的变量通过cahnnel来请求的goroutine
func teller() {
	//流水线的每一个阶段都能够避免在将变量传送到下一阶段时再去访问它
	var balance1 int //这个变量对所有访问都是线性的

	//串行绑定:这个变量会被绑定到流水线的一个阶段，传送完后被绑定到下一个，以此类推
	for {
		select {
		case amout := <-deposits:
			balance1 += amout
		case balances <- balance1:
		}
	}
}
func bank1() {
	go teller() //启动监控协程
	go func() {
		Deposit1(10)
		fmt.Println(Balance1())
	}()
	go func() {
		Deposit1(100)
		fmt.Println(Balance1())
	}()
	go func() {
		Deposit1(1000)
		fmt.Println(Balance1())
	}()
	//一个一个协程的处理 并发意义何在？ 就为了共享变量
	for {
	}
}

//互斥
//允许很多goroutine去访问变量，但是在同一个时刻最多只有一个goroutine在访问
//确保共享变量在程序执行时的关键点上能够保证不变性
var (
	sema     = make(chan struct{}, 1)
	balance2 int
)

//协程调用Deposit2上锁，其他协程调用Balance2和Deposit2函数直接阻塞，等待解锁
func Deposit2(amout int) {
	sema <- struct{}{} //向缓冲为1的channel写数据：上锁,其他协程调用此函数直接阻塞
	balance2 += amout
	<-sema //函数调用完成后解锁
}

//协程调用Balance2上锁，其他协程调用Balance2和Deposit2函数直接阻塞，等待解锁
func Balance2() int {
	sema <- struct{}{} //若上面的函数被锁,调用此函数的协程阻塞在这
	b := balance2
	<-sema
	return b
}
func bank2() {
	go func() {
		Deposit2(10)
		fmt.Println(Balance2())
	}()
	go func() {
		Deposit2(100)
		fmt.Println(Balance2())
	}()
	go func() {
		Deposit2(1000)
		fmt.Println(Balance2())
	}()
	for {
	}

}

//sync.Mutex 和上面一样，原理就是那个原理
var (
	O        sync.Mutex
	balance3 int
)

func Deposit3(amout int) {
	O.Lock()
	balance3 += amout
	O.Unlock() //在Lock和Unlock之间的代码段中的内容goroutine可以随便读取或者修改，这个代码段叫做临界区
}
func Balance3() int {
	O.Lock()
	b := balance3
	O.Unlock() //多和defer配合使用
	return b
}
func Balance4() int {
	O.Lock()
	defer O.Unlock()
	return balance3
}
func Withdraw(amout int) bool {
	O.Lock()
	defer O.Unlock()
	//Deposit3(-amout)  无法调的，因为已经上锁了 可以将要调用的函数分离
	Deposit0(-amout) //注意了,是Deposit0
	if Balance0() < 0 {
		Deposit0(amout)
		return false
	}
	return true

}
func Deposit0(amout int) {
	balance3 += amout
}
func Balance0() int {
	return balance3
}

//读写锁
//多读单写:允许多个只读操作并行执行，但写操作会完全互斥
var OO sync.RWMutex

func Balance5() int {
	OO.RLock()
	defer OO.RUnlock()
	return balance3
}
func monitor() {
	tick := time.NewTicker(1 * time.Second)
	defer tick.Stop()
	for {
		select {
		case <-tick.C:
			fmt.Println(Balance5())
		}
	}
}

//用户输入任意字符存钱，并不会影响查询
func bank3() {
	go monitor()
	go func() {
		for {
			os.Stdin.Read(make([]byte, 1)) //等待用户输入一个字符
			Deposit3(100)
		}

	}()
	time.Sleep(20 * time.Second)
} //sync.Once
//核心思想是使用原子计数记录被执行的次数。使用Mutex Lock Unlock锁定被执行函数，防止被重复执行。
func test15() {
	var once sync.Once
	onceBody := func() {
		fmt.Println("Only Once")
	}
	done := make(chan bool)
	for i := 0; i < 5; i++ {
		j := i
		go func(int) {
			once.Do(onceBody) //once.do(函数名)  函数只执行一次！执行一次后，其他协程不执行了
			fmt.Println(j)
			done <- true
		}(j)
	}
	<-done
	time.Sleep(2 * time.Second)
}
