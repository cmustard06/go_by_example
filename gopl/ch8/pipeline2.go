package main

import "fmt"

//对于管道close的简单实用

func main() {
	naturals := make(chan int)
	squares := make(chan int)

	//counter，上线为100.到了就关闭管道
	go func() {
		for x := 0; x < 100; x++ {
			naturals <- x
		}
		close(naturals) //对于close，只能在管道的发送方调用，接收方不能调用这个close，会导致宕机
	}()

	//square
	go func() {
		for x := range naturals { //range语法糖，会自动判断管道是否被关闭
			squares <- x * x
		}
		//如果执行完了，在这边关闭squares通道
		close(squares)
	}()

	//print
	//如果我们不使用range来判断管道是否关闭，我们也可以使用如下方式
	for {
		r, ok := <-squares //可以使用这种方式的表达式来判断管道是否关闭
		if !ok {
			//管道关闭了
			break
		}
		fmt.Println(r)
	}
}
