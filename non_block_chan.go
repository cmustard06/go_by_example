package main

import "fmt"

func main() {
	messages := make(chan string)
	signals := make(chan bool)
	select {
	case msg := <-messages:
		fmt.Println("rec message:", msg)
	default:
		fmt.Println("no message received")
	}


	msg := "hi"
	select {
	case messages <- msg:  //无法将信息发送到管道中，因为管道既没有缓冲区，也没有接受者
		fmt.Println("sent message:", msg)
	default:
		fmt.Println("no message send")
	}


	//尝试使用非阻塞的形式获取管道中的值
	select {
	case msg := <-messages:
		fmt.Println("received message:", msg)
	case sig := <-signals:
		fmt.Println("received signal:", sig)
	default:
		fmt.Println("no activity")
	}
}
