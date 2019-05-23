package main

import (
	"fmt"
	"os"
	"time"
)

//该select语句使每一次迭代使用1s来等待终止
/*
tick函数很方便使用，但是他仅仅在应用的整个生命周期都需要时才合适
否则我们需要使用如下模式
ticker := time.NewTicker(1*time.Second)
<- ticker.C //从ticker接收
ticker.Stop() //造成ticker的goroutine终止
*/

func main() {
	abort := make(chan struct{})

	go func() {
		os.Stdin.Read(make([]byte, 1))
		abort <- struct{}{}
	}()

	fmt.Println("开始倒计时。按返回中止")
	//tick := time.Tick(1 * time.Second)
	//for countdown := 10; countdown > 0; countdown-- {
	//	fmt.Println(countdown)
	//	select {
	//	/*
	//	time.Tick函数的行为很像创建一个goroutine在循环里面调用
	//	time.Sleep，然后在他每次醒来发送一个事件，
	//	当上面的倒计时函数返回时，它停止从tick中接收事件，
	//	但是计时器goroutine还在继续运行，徒劳的向一个没有goroutine
	//	在接收的通道不断发送，发生goroutine泄漏
	//	*/
	//	case <-tick:
	//		//不执行任何操作
	//	case <-abort:
	//		fmt.Println("启动被终止")  //这里tick没有被终止，管道一直存在
	//		return
	//	}
	//}
	ticker := time.NewTicker(1 * time.Second)

	for countdown := 10; countdown > 0; countdown-- {
		fmt.Println(countdown)
		select {
		/*
			time.Tick函数的行为很像创建一个goroutine在循环里面调用
			time.Sleep，然后在他每次醒来发送一个事件，
			当上面的倒计时函数返回时，它停止从tick中接收事件，
			但是计时器goroutine还在继续运行，徒劳的向一个没有goroutine
			在接收的通道不断发送，发生goroutine泄漏
		*/
		case <-ticker.C:
			//不执行任何操作
		case <-abort:
			fmt.Println("启动被终止")
			//终止ticker管道
			ticker.Stop()
			return
		}
	}
	launch()
}

func launch() {
	fmt.Println("Lift off!")
}
