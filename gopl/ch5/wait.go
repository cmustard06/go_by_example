package main

import (
	"time"
	"net/http"
	"log"
	"fmt"
	"os"
)

//在某一段时间内循环尝试某一动作

func WaitForServer(url string) error{
	const timeout = 1*time.Minute  //循环尝试时间
	deadline := time.Now().Add(timeout)
	for tries := 0;time.Now().Before(deadline);tries++{  //开始尝试
		resp, err := http.Get(url)
		if err==nil{
			resp.Body.Close()
			return nil  //成功
		}
		log.Printf("server not responding(%s); retrying...",err)
		time.Sleep(time.Second<<uint(tries)) //设置每次尝试的间隔，指数级增长
	}
	return fmt.Errorf("server %s failed to respond after %s", url, timeout)
}

func main(){
	url := "http://127.0.0.1:8000"
	if err := WaitForServer(url); err != nil {
		fmt.Fprintf(os.Stderr, "Site is down: %v\n", err)
		os.Exit(1)
	}
}