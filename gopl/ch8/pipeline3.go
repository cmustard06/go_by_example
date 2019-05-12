package main

import "fmt"

//单向通道的使用

func counter(natuals chan<- int) { //这个函数中的natuals只允许发送操作
	for x := 0; x < 100; x++ {
		natuals <- x
	}
	close(natuals)
}

func square(squares chan<- int, natuals <-chan int) {
	for x := range natuals {
		squares <- x * x
	}
	close(squares)
}

func main() {
	natuals := make(chan int) //channel是引用类型，类似于map，函数之间的传递都是引用传递，可以进行比较，相同说明用到同一个管道
	squares := make(chan int)

	go counter(natuals) //这里会将natuals隐式转换为单向通道，主意只能双向通道能转，反之不行
	go square(squares, natuals)
	for {
		r, ok := <-squares
		if !ok {
			break
		}
		fmt.Println(r)
	}
}
