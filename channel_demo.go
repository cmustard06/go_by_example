package main

import (
	"fmt"
	"time"
)

func worker(done chan bool){
	fmt.Println("working...")
	time.Sleep(2*time.Second)
	fmt.Println("done")
	done <- true
}

func ping(pings chan<- string, msg string){   //这个管道只接受value
	pings <- msg
}

func pong(pings <- chan string, pongs chan<-string){
	msg := <- pings
	pongs <- msg
}

func main(){
	//done := make(chan bool,1)
	//go worker(done)
	//
	//<- done

	pings := make(chan string, 1)
	pongs := make(chan string, 1)
	ping(pings, "passwd")
	pong(pings, pongs)
	fmt.Println(<-pongs)
}
