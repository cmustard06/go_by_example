package main

import (
	"fmt"
	"io"
	"log"
	"net"
	"os"
)

//无缓冲通道的使用
/*
对于无缓冲通道上的发送的操作将会阻塞，直到另一个goroutine读取这个管道中的数据，这就说明值传递完成，两个goroutine都可以继续
工作了

该程序只是简单的示例 mac 使用ctrl+D发送EOF信号
*/

func mustCopy(dst io.Writer, src io.Reader) {
	//当用户关闭输入流的时候(mac 使用ctrl+D)，该函数会返回，
	if _, err := io.Copy(dst, src); err != nil {
		log.Fatal(err)
	}
}

func main() {
	conn, err := net.Dial("tcp", "localhost:8000")
	if err != nil {
		log.Fatal(err)
	}
	done := make(chan struct{})
	//这里单独起一个协程去处理服务器发过来的数据
	go func() {
		io.Copy(os.Stdout, conn)
		log.Println("done")
		done <- struct{}{} //给主线程一个信号,这里通常习惯用int或者bool型
	}()
	mustCopy(conn, os.Stdin)
	conn.Close()
	fmt.Println("coming to end")
	<-done //如果管道中没有数据，这里一直会阻塞
	fmt.Println("finish!!!")

}
