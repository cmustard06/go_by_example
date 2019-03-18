package main

import (
	"os"
	"os/signal"
	"syscall"
	"fmt"
)

func main(){
	sigs := make(chan os.Signal, 1)
	done := make(chan bool,1)

	//接受指定信号的通知
	signal.Notify(sigs,syscall.SIGINT, syscall.SIGTERM)

	go func() {
		sig := <-sigs
		fmt.Println()
		fmt.Println(sig)
		done <- true
	}()

	fmt.Println("awaitinh signal")
	<-done
	fmt.Println("exiting")
}
