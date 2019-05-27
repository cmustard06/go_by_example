package main

import (
	"sync"
	"fmt"
)

//利用互斥锁来重新实现2中的

var(
	mu sync.Mutex  //互斥锁
	balance int
)

func Deposit(amount int){
	mu.Lock()
	balance = balance+amount
	mu.Unlock()
}

func Balance() int{
	mu.Lock()
	b := balance
	mu.Unlock()
	return b
}

func main(){
	var n sync.WaitGroup
	for i := 1; i <= 1000; i++ {
		n.Add(1)
		go func(amount int) {
			Deposit(amount)
			n.Done()
		}(i)
	}
	n.Wait()

	if got, want := Balance(), (1000+1)*1000/2; got != want {
		fmt.Errorf("Balance = %d, want %d", got, want)
	}
}
