package main

import (
	"net"
	"fmt"
	"bufio"
	"time"
	"io"
	"os"
)

func client(){
	conn, e := net.DialTimeout("tcp", "192.168.13.180:70", time.Second)
	if e!=nil{
		panic(e)
	}
	fmt.Fprint(conn,"GET / HTTP/1.0\r\n\r\n")

	result := make([]byte,1000)
	length, err := bufio.NewReader(conn).Read(result)
	if err!=nil{
		panic(err)
	}
	fmt.Println(length)
	fmt.Println(string(result))
}

func handConnection(conn net.Conn){
	requestData := make([]byte, 1000)
	length, err := bufio.NewReader(conn).Read(requestData)
	if err!=nil && err!=io.EOF{
		panic(err)
	}
	fmt.Println(length, string(requestData))
	//回复
	n, err := conn.Write([]byte("HTTP/2.0 200 OK\r\n\r\nhello"))
	if err!=nil{
		fmt.Println(err)
	}

	fmt.Println("send:",n)
	conn.Close()
}

func server(){
	fmt.Println("starting listen...")
	sock, err := net.Listen("tcp", "127.0.0.1:8080")
	if err!=nil{
		fmt.Println(err)
	}
	for{
		conn, err := sock.Accept()
		if err!=nil{
			panic(err)
		}
		go handConnection(conn)
	}
}

func main(){
	//client()
	server()
	os.Exit(0)
	p := fmt.Println
	p(net.IPv4len)
	p(net.JoinHostPort("127.0.0.1","8888")) //127.0.0.1:8888

	//解析地址,反向解析
	names,err := net.LookupAddr("8.8.8.8") //[google-public-dns-a.google.com.]
	if err!=nil{
		fmt.Println(err)
	}
	fmt.Println(names)
	//获取host cname
	cname,_ := net.LookupCNAME("www.google.com")
	fmt.Println(cname)

	//dns解析
	addr,_ := net.LookupHost("merak.cc")
	fmt.Println(addr) //[45.78.21.125]

	//查找服务端口
	port,_ := net.LookupPort("tcp","ssh")
	p(port) //22

	//CIDR解析
	ipv4Addr,ipv4Net,_ := net.ParseCIDR("192.0.2.0/16")
	p(ipv4Addr,ipv4Net) //192.0.2.0 192.0.2.0/24

	//比较ip是否相等
	p(net.ParseIP("192.168.1.1").Equal(net.IPv4(192,168,1,1))) //true

}
