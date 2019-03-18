package main

import (
	"sync/atomic"
	"time"
	"fmt"
)

func main(){
	var ops uint64

	for i:=0;i<500;i++{
		go func(){
			// 实现原子操作，保证结果正确
			atomic.AddUint64(&ops,1)
			//ops = ops+1
			time.Sleep(time.Millisecond)
		}()
	}

	time.Sleep(time.Second)

	opsFinal := atomic.LoadUint64(&ops)
	fmt.Println("ops:",opsFinal)
}
