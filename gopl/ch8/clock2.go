package main

import (
	"net"
	"io"
	"time"
	"log"
)

//每秒向客户端更新时间

func handleConn(c net.Conn){
	defer c.Close()
	for{
		_, err := io.WriteString(c, time.Now().Format("15:04:05\n")) //需要记住格式化 字符串是1月2日下午3点4分5秒零六年UTC­0700
		if err!=nil{
			return //e.g 客户端主动断开连接
		}
		time.Sleep(1*time.Second)

	}
}

func main(){
	listener, err := net.Listen("tcp", "localhost:8000")
	if err!=nil{
		log.Fatal(err)
	}
	for{
		conn, err := listener.Accept()
		if err!=nil{
			log.Print(err) // e.g 连接被终止
			continue
		}
		go handleConn(conn)
	}
}

