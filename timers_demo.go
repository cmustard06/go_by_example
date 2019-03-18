package main

import (
	"time"
	"fmt"
)

func main(){
	timer1 := time.NewTimer(2* time.Second)
	current := time.Now()
	<- timer1.C  //阻塞，直到发送一个定时器到期的值
	fmt.Println("Timer 1 expired", time.Now().Sub(current))

	timer2 := time.NewTimer(time.Second)

	go func(){
		<- timer2.C //阻塞
		fmt.Println("Timer 2 expired")
	}()
	//time.Sleep(3*time.Second)
	stop2 := timer2.Stop()
	if stop2{
		fmt.Println("Timer 2 stopped")
	}
}
