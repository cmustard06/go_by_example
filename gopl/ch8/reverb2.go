package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"strings"
	"time"
)

func handleConn(c net.Conn) {
	//处理接收到的数据
	defer c.Close()
	scanner := bufio.NewScanner(c)
	for scanner.Scan() {
		go echo(c, scanner.Text(), 1*time.Second) //这样客户端如果同时发送过来消息，消息就不会出现延迟了
	}
}

func echo(c net.Conn, shout string, delay time.Duration) {
	fmt.Fprintf(c, "\n%s", strings.ToUpper(shout))
	time.Sleep(delay)
	fmt.Fprintf(c, "\n%s", shout)
	time.Sleep(delay)
	fmt.Fprintf(c, "\n%s", strings.ToTitle(shout))
}

func main() {
	client, err := net.Listen("tcp", "localhost:8000")
	fmt.Println("Listening 8000")
	if err != nil {
		log.Fatal(err)
	}
	for {
		conn, err := client.Accept()
		if err != nil {
			log.Println(err)
			continue
		}
		go handleConn(conn)
	}
}
