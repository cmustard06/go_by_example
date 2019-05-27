package main

import (
	"net"
	"log"
	"fmt"
	"bufio"
)

//简单的聊天服务器

type client chan string //一个类型

var (
	entering = make(chan client) // chan chan  string,entering channel用于读取和写入一个chan string channel 类型的channel
	leaving  = make(chan client)
	messages = make(chan string) //所有来自客户端的信息
)

func main() {
	listener, err := net.Listen("tcp", "localhost:8000")
	if err != nil {
		log.Fatal(err)
	}
	//客户端广播
	go broadcaster()
	//获取来自client的连接
	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Print(err)
			continue
		}
		//启动一个协程去处理
		go handleConn(conn)
	}
}

func broadcaster() {
	clients := make(map[client]bool)
	for {
		select {
		case cli := <-entering:  //取出刚建立连接的channel
			clients[cli] = true
		case msg := <-messages:
			//将传入的消息广播到所有的传出消息
			for cli := range clients {
				cli <- msg //将数据传入client管道
			}
		case cli := <-leaving: //取出完成任务的channel
			delete(clients, cli)
			close(cli)
		}
	}
}

func handleConn(conn net.Conn) {
	ch := make(chan string) //传给客户端的信息
	go clientWriter(conn, ch)

	//获取客户端相关信息
	addr := conn.RemoteAddr().String()
	fmt.Println(addr)
	entering <- ch
	messages <- addr + "has arrived\n"
	input := bufio.NewScanner(conn)
	for input.Scan() {
		messages <- addr + ": " + input.Text()
	}
	messages <- addr + " has left\n"
	leaving <- ch
	conn.Close()
}

func clientWriter(conn net.Conn, ch <-chan string) {
	for msg := range ch {
		fmt.Fprintf(conn, msg)
	}
}
