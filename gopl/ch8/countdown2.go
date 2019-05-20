package main

import (
	"fmt"
	"os"
	"time"
)

//在版本1的基础上，增加从命令行取消发射的功能
//这里就需要两个goroutine，一个用于发射，一个
//用于监听取消发射功能，这里可引入select{}多路复用

func main() {
	//创建一个终止管道
	abort := make(chan struct{})
	//启动一个goroutine时刻关注来自命令行的回车键
	go func() {
		os.Stdin.Read(make([]byte, 1)) //等待读取一个字节
		abort <- struct{}{}
	}()
	fmt.Println("开始倒计时。按回车键即可终止发射")
	/*
		select一直等待，直到一次通信来告知有一些情况可以执行，
		然后就执行相应的通信，而其他情况所对应的语句将不会被执行
		，对于没有对应情况的select，将会一直等待
	*/
	select {
	case <-time.After(10 * time.Second):
		fmt.Println("....")
		//nothing
	case <-abort:
		fmt.Println("发射被终止")
		return

	}
	launch()
}

func launch() {
	fmt.Println("起飞")
}
