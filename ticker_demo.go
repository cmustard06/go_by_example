package main

import (
	"time"
	"fmt"
)

func main(){
	ticker := time.NewTicker(500*time.Millisecond)

	go func() {
		for t:= range ticker.C{    // 可以每隔多少时间执行一次操作
			fmt.Println("Tick at ",t)
		}
	}()

	time.Sleep(1600*time.Millisecond)
	//ticker.Stop()
	fmt.Println("Ticker stopped")
}
