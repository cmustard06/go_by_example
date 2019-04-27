package main

import (
	"time"
	"log"
)

//通过defer关键字追踪函数执行和退出时候的状态

func bigSlowOperation(){
	defer trace("bigSlowOperation")()  //通过defer为函数增加进入和结束，函数刚执行的时候，执行trace()，函数退出执行trace()()
	time.Sleep(10*time.Second)
}

func trace(msg string)func(){
	start := time.Now()
	log.Printf("enter %s",msg)
	return func() {
		log.Printf("exit %s (%s)",msg,time.Since(start))
	}
}

func main(){
	bigSlowOperation()
}
