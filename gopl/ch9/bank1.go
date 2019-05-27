package main

import (
	"fmt"
	"log"
)

// 模拟银行账单操作，由于并发不安全，导致读取存款的时候，读取的结果可能是没有加上之前的数据

var deposits = make(chan int) //发送金额存款
var balances = make(chan int) //收到，余额

func Deposit(amount int) {
	deposits <- amount
}

func Balance() int {
	return <-balances
}

func teller() {
	var balance int
	for {
		select {
		case amount := <-deposits:
			balance += amount
		case balances <- balance:
		}

	}
}

func init(){
	go teller()
}

func main(){
	done := make(chan struct{})

	//A同学存款
	for i:=0;i<1000;i++{
		go func() {
			Deposit(200)
			fmt.Println("=",Balance())
			done <- struct{}{}
		}()
	}
	for i:=0;i<1000;i++{
		<-done
	}



	if got,want := Balance(),200*1000;got!=want{
		log.Printf("Balance=%d want %d",got,want)
	}

}