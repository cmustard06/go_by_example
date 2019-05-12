package main

import "fmt"

//管道的简单使用，实际中不会像这个程序这样写,这个程序会一直执行下去

func main() {
	naturals := make(chan int)
	squares := make(chan int)

	//计数器
	go func() {
		for x := 0; ; x++ {
			naturals <- x //将当前的计数值传入管道中，如果另一个协程不取，就阻塞
		}
	}()

	//计算平方
	go func() {
		for {
			x := <-naturals  //取出x
			squares <- x * x //发送到管道中
		}
	}()

	//主goroutine中打印
	for {
		fmt.Println(<-squares)
	}
}
