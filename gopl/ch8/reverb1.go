package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"strings"
	"time"
)

//构建一个每一个连接使用多个goroutine来处理的echo服务器

func main() {
	listener, err := net.Listen("tcp", "localhost:8000")
	if err != nil {
		log.Fatal(err)
	}

	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Println(err)
			continue
		}
		go handleConn(conn)
	}

}

func handleConn(c net.Conn) {
	scanner := bufio.NewScanner(c)
	for scanner.Scan() {
		//echo
		echo(c, scanner.Text(), 1*time.Second)
	}
	c.Close()

}

func echo(c net.Conn, shout string, delay time.Duration) {
	fmt.Fprintf(c, "\t%s", strings.ToUpper(shout))
	time.Sleep(delay)
	fmt.Fprintf(c, "\t%s", shout)
	time.Sleep(delay)
	fmt.Fprintf(c, "\t%s", strings.ToLower(shout))
}
