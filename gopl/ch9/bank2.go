package main

import (
	"sync"
	"fmt"
)

//提供并发安全的

var (
	sema = make(chan struct{},1) //信号量
	balance int
)

func Deposit(amount int) {
	sema <- struct{}{}
	balance = balance+amount
	<- sema //释放
}

func Balance()int{
	sema <- struct{}{}
	b := balance //读取这个全局变量时，其他协程不能操作这个变量
	<-sema
	return b
}


func main(){
	var n sync.WaitGroup
	for i:=0;i<=10000;i++{
		n.Add(1)
		go func(amount int) {
			Deposit(amount)
			n.Done()
		}(i)
	}
	n.Wait()
	if got, want := Balance(), (10000+1)*10000/2; got != want {
		fmt.Printf("Balance = %d, want %d", got, want)
	}
}