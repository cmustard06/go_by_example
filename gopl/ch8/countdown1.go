package main

import (
	"fmt"
	"time"
)

//模拟火箭发射倒计时，

func main() {
	fmt.Println("开始倒计时")
	tick := time.Tick(1 * time.Second) // 该函数返回一个通道，定期(这里是1s)向通道里面发送事件
	for countdown := 10; countdown > 0; countdown-- {
		fmt.Printf("\t %d\n", countdown)
		fmt.Print(<-tick) //从channel中接收事件
	}
	launch()
}

func launch() {
	fmt.Println("\t起飞")
}
