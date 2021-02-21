package main

import (
	"context"
	"fmt"
	"os"
	"runtime"
	"time"
)

//创建协程goroutine
func test01() {
	go func() { //创建协程
		for {
			fmt.Println("this is a newTask")
			time.Sleep(time.Second) //延时1s
		}
	}()

	for i := 0; i < 5; i++ { //当主协程退出了，子协程也退出了(有可能子协程还没来得及调用，主协程就退出了)
		fmt.Println("this is a main goroutine")
		time.Sleep(time.Second) //延时1s
	}
}

//runtime包
func test02() {
	go func() {
		for i := 0; i < 5; i++ {
			fmt.Println("ohhh")
		}
	}()
	//让出时间片，先让别的协程执行，他执行完，再回来执行此协程
	runtime.Gosched()

	for i := 0; i < 2; i++ {
		fmt.Println("aaaaa")
	}
}
func test02_1() {
	go func() {
		fmt.Println("1")
		func() {
			defer fmt.Println("2")
			runtime.Goexit() //立即终止所在协程,defer可以被执行
			fmt.Println("4")
		}()
		fmt.Println("3")
	}()
	for {
	}
}
func test02_2() {
	n := runtime.GOMAXPROCS(1) //指定单核运算
	//	n := runtime.GOMAXPROCS(4)
	fmt.Println("n=", n)
	//指定多核看下效果
	for {
		go fmt.Print(1)
		fmt.Print(0)
	}

}

//多任务资源竞争
//通过channel实现同步
//两任务同时进行,任务1先执行打印,任务2从管道读数据(管道没有数据)阻塞，直到任务1执行完打印操作后往管道里写数据，任务2关闭阻塞开始打印
func test03() {
	var ch = make(chan int)
	go func() {
		test03_1("hello")
		ch <- 1000 //给管道写数据
	}()
	go func() {
		<-ch //读数据，若管道中没有数据 堵塞
		test03_1("world")
	}()
	for {

	}
}
func test03_1(str string) {
	for _, data := range str {
		fmt.Printf("%c", data)
		time.Sleep(time.Second)
	}
}
func test04() {
	ch := make(chan string)
	defer fmt.Println("main is ok")
	go func() {
		for i := 0; i < 2; i++ {
			fmt.Println(i)
			time.Sleep(time.Second)
			ch <- "ooh"
		}
	}()
	str := <-ch
	fmt.Println(str)
}

//无缓冲channel
//子协程打印并向管道写数据(主协程堵塞)，主协程读数据并打印(子协程阻塞),子协程打印并向管道写数据,此时主协程来不及读数据打印，即子协程打印了两次(但只写了一次),主协程读并打印，子协程写，主协程读并打印
/*子协程： 0*/
/*主协程:  0*/
/*子协程： 1*/
/*子协程： 2*/
/*主协程:  1*/
/*主协程:  2*/
func test05() {
	ch := make(chan int)
	go func() {
		for i := 0; i < 3; i++ {
			fmt.Println("子协程：", i)
			ch <- i
		}
	}()
	time.Sleep(2 * time.Second)
	for i := 0; i < 3; i++ {
		n := <-ch
		fmt.Println("主协程:", n)
	}
}

//有缓冲的channel
func test06() {
	ch := make(chan int, 3)
	go func() {
		for i := 0; i < 10; i++ {
			ch <- i
			fmt.Println("子协程:", i)
		}
	}()
	time.Sleep(2 * time.Second)
	for i := 0; i < 10; i++ {
		n := <-ch
		fmt.Println("主协程:", n)
	}

}
func test07() {
	ch := make(chan int)
	go func() {
		for i := 0; i < 5; i++ {
			ch <- i
		}
		close(ch) //关闭管道，关闭后不能再向其写数据
	}()
	for {
		if n, ok := <-ch; ok == true {
			fmt.Println("n=", n)
		} else {
			break
		}
	}
}

//遍历管道数据
func test07_1() {
	ch := make(chan int)
	go func() {
		for i := 0; i < 5; i++ {
			ch <- i
		}
		close(ch) //关闭管道，关闭后不能再向其写数据
	}()

	for n := range ch {
		fmt.Println("n=", n)
	}
}
func test8() {
	ch := make(chan int)
	//双向channel隐式转换为单向channel(单向不能转双向)
	var writech chan<- int = ch //只写不能读
	writech <- 100
	//	<-writech
	//(receive from send-only type chan<- int)
	var readch <-chan int = ch //只读不能写
	//readch <- 100
	//(send to receive-only type <-chan int)
	<-readch
}
func test08_1() {
	ch := make(chan int)
	go producer(ch)
	consumer1(ch)
}
func producer(out chan<- int) {
	for i := 0; i < 10; i++ {
		out <- i * i
	}
	close(out)
}
func consumer(in <-chan int) {
	for n := range in {
		fmt.Println("n=", n)
	}
}

//当生产者停止时，只是停止向channel写并没有关闭channel因为要防止其他协程读取该channel发送错误，这时其他协程去读该channel全为0
func consumer2(in <-chan int) {
	for {
		select {
		case s, _ := <-in:
			fmt.Println("n=", s)
		default:
		}

	}
}

func consumer1(in <-chan int) {
	for {
		select {
		case s, ok := <-in:
			if ok { //判断channel中还有没有数据，若没有返回
				fmt.Println("n=", s)
			} else {
				return
			}
		default:
		}

	}
}

//定时器 timer相应一次
func test09() {
	timer := time.NewTimer(2 * time.Second)
	fmt.Println(time.Now())
	t := <-timer.C
	fmt.Println(t)
	//延时
	timer1 := time.NewTimer(2 * time.Second)
	<-timer1.C //堵塞直到timer1有数据了
	fmt.Println("时间到")

	time.Sleep(2 * time.Second)
	fmt.Println("时间到")
	<-time.After(2 * time.Second) //定时2s,阻塞2s
	fmt.Println("时间到")

	timer2 := time.NewTimer(3 * time.Second)
	go func() {
		<-timer2.C
		fmt.Println("子协程可以打印了，时间到了")
	}()
	//停止定时器
	timer2.Stop()
	for {

	}
}
func test09_1() { //timer只会相应一次
	timer := time.NewTimer(5 * time.Second)
	//重置定时器
	timer.Reset(1 * time.Second)
	<-timer.C
	fmt.Println("时间到")

}

//Ticker的使用
func test10() {
	ticker := time.NewTicker(1 * time.Second)
	i := 0
	for {
		n := <-ticker.C //每隔时间段回向管道发送当前时间
		fmt.Println("n", n)
		i++
		fmt.Println("i=", i)
		if i == 5 {
			ticker.Stop()
			break
		}
	}
}

//select监听channel的数据流动

//select实现超时机制
func test11() {
	ch := make(chan int)
	quit := make(chan bool)
	go func() {
		for {
			select {
			case n := <-ch:
				fmt.Println("n=", n)
			case <-time.After(3 * time.Second):
				fmt.Println("超时")
				quit <- true //关闭阻塞
			}
		}
	}()
	for i := 0; i < 5; i++ {
		ch <- i
	}
	<-quit //阻塞
	fmt.Println("结束")
}

//现在每一次计数循环的迭代都需要等待两个channel中的其中一个返回事件了：ticker channel当一切正常时或者异常时返回的abort事件。我们无法做到从每一个channel中接收信息，如果我们这么做的话，如果第一个channel中没有事件发过来那么程序就会立刻被阻塞，这样我们就无法收到第二个channel中发过来的事件
func test11_1() {
	fmt.Println("START")
	tick := time.Tick(1 * time.Second)
	for s := 10; s > 0; s-- {
		fmt.Println("ohhhhh")
		<-tick
	}
	launch()
	abort := make(chan struct{})
	go func() {
		os.Stdin.Read(make([]byte, 1)) //等待用户输入一个字符
		abort <- struct{}{}
	}()
}
func launch() {
	fmt.Println("起飞")
}
func test11_2() {
	abort := make(chan struct{})
	go func() {
		os.Stdin.Read(make([]byte, 1)) //等待用户输入一个字符
		abort <- struct{}{}
	}()
	select {
	case <-time.After(10 * time.Second):
		launch()
	case <-abort:
		fmt.Println("有内鬼，终止交易")
		return
	}
}

//串联的Channels(Pipeline) 一个channels的输出作为下一个channels的输入
func test12() {
	ch1, ch2 := make(chan int), make(chan int)
	go func() {
		time.Sleep(1 * time.Second)
		for i := 1; i < 10; i++ {
			ch1 <- i

		}
		close(ch1)
		fmt.Println("ch1已停止")
	}()
	go func() {
		for {
			time.Sleep(1 * time.Second)
			x, ok := <-ch1
			if !ok {
				break //ch1没数据了
			}

			ch2 <- x
		}
		close(ch2)
	}()
	for {
		time.Sleep(1 * time.Second)
		fmt.Println(<-ch2)

		if <-ch2 == 0 {
			break
		}
	}
}

//range迭代channels
func test12_1() {
	ch1, ch2 := make(chan int), make(chan int)
	go func() {
		for x := 0; x < 100; x++ {
			ch1 <- x
		}
		close(ch1)
	}()
	go func() {
		for x := range ch1 {
			ch2 <- x
		}
		close(ch2)
	}()
	for x := range ch2 {
		fmt.Println(x)
	}
}
func test13() {
	ch := make(chan int, 1)
	for i := 0; i < 10; i++ {
		select {
		case x := <-ch:
			fmt.Println(x)
		case ch <- i:
		}
	}
}

//循环结束后，tick停止接受事件，但ticker依然存活且仍不断向tick了写入数据---goroutine泄露
//goroutine内存泄露
//go中的内存泄露一般都是goroutine泄露，就是goroutine没有被关闭，或者没有添加超时控制，让goroutine一只处于阻塞状态，不能被GC。
func test13_1() {
	abort := make(chan struct{})
	go func() {
		os.Stdin.Read(make([]byte, 1)) //等待用户输入一个字符
		abort <- struct{}{}
	}()
	tick := time.Tick(1 * time.Second) //Tick是NewTicker的封装，只提供对Ticker的channel访问，若不需要关闭Ticker，就可用此函数
	fmt.Println("START")
	for s := 10; s > 0; s-- {
		fmt.Println(s)
		select {
		case <-tick: //注意这里是直接tick而不是tick.C
		case <-abort:
			fmt.Println("有内鬼，终止交易")
			return
		}

	}
	launch()
}
func test13_2() {
	abort := make(chan struct{})
	go func() {
		os.Stdin.Read(make([]byte, 1)) //等待用户输入一个字符
		abort <- struct{}{}            //这种方式只能关闭一个协程，并不试用并发退出的情况,可以通过广播的方式退出并发(参考mydu中的channel done)
	}()
	fmt.Println("START")
	for { //轮询channel
		select {
		case <-abort:
			fmt.Println("有内鬼，终止交易")
			return
		default:

		}
	}
}

//Ticker的关闭
func test14() {
	ticker1 := UseTickerWrong()
	time.Sleep(20 * time.Second)
	ticker1.Stop() //Sopt确实会停止Ticker，Ticker不会再写，但是他并不会关闭channel，防止读取channel发生错误(但是读channel的协程会被堵塞)
}

func UseTickerWrong() *time.Ticker {
	ticker := time.NewTicker(5 * time.Second)
	go func(ticker *time.Ticker) {
		for range ticker.C {
			fmt.Println("Ticker1....")
		} //关闭ticker后，会堵塞在这里

		fmt.Println("Ticker1 Stop")
	}(ticker)

	return ticker
}

//正确的关闭Ticker:使用select监测ticker是否关闭，若关闭就返回
var done = make(chan struct{})

func UseTicker() chan bool {
	ticker := time.NewTicker(1 * time.Second)
	stopChan := make(chan bool)
	go func() { //输入任意键关闭done
		os.Stdin.Read(make([]byte, 1))
		done <- struct{}{}
	}()
	go func(ticker *time.Ticker) {
		defer ticker.Stop()
		for {
			select {
			case <-ticker.C:
				fmt.Println("Ticker2...")
			case <-done:
				fmt.Println("Ticker2 stop")
				return
			case stop := <-stopChan: //当stop通道读到true的时候，函数返回
				if stop {
					fmt.Println("Ticker2 stop")
					return
				}
			}
		}
	}(ticker)
	return stopChan
}
func test14_1() {
	ch := UseTicker()
	time.Sleep(10 * time.Second)
	if test14_01() {
		ch <- true
	}
	ch <- true
	close(ch)
	fmt.Println("成功退出")

}

func test14_01() bool { //轮询函数
	select {
	case <-done:
		return true
	default:
		return false
	}
}

//context:当一个goroutine在衍生一个goroutine时，context可以跟踪到子goroutine，从而达到控制他们的目的
func test155() {
	//WithCancel创建一个可取消的子Context
	ctx, cancel := context.WithCancel(context.Background()) //Background返回一个空的context，用于整个Context树的根节点

	go func(ctx context.Context) { //Context 跟踪goroutin
		for {
			select {
			case <-ctx.Done():
				fmt.Println("JOB is DONE")
				return

			default:
				fmt.Println("JOB is RUN")
				time.Sleep(time.Second * 2)
			}
		}
	}(ctx)
	time.Sleep(time.Second * 10)
	fmt.Println("JOB need STOP!!!")
	cancel()
	time.Sleep(time.Second * 3)
}
func watch(ctx context.Context, name string) {
	for {
		select {
		case <-ctx.Done():
			fmt.Println(name, " is DONE")
			return
		default:
			fmt.Println(name, " is RUN")
			time.Sleep(2 * time.Second)
		}

	}

}

//它就像一个控制器一样，按下开关后，所有基于这个 Context 或者衍生的子 Context 都会收到通知，这时就可以进行清理操作了，最终释放 goroutine
func test156() {
	ctx, cancel := context.WithCancel(context.Background()) //Background返回一个空的context，用于整个Context树的根节点
	go watch(ctx, "JOB1")
	go watch(ctx, "JOB2")
	go watch(ctx, "JOB3")
	time.Sleep(time.Second * 10)
	fmt.Println("JOB need STOP!!!")
	cancel()
	time.Sleep(time.Second * 5)

}
func test157() {
	ctx := context.WithValue(context.Background(), 1, "root")
	fmt.Println(ctx.Value(1))
}
