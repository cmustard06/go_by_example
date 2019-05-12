package main

import (
	"crypto/tls"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"runtime"
	"runtime/debug"
	"runtime/pprof"
	"time"
)

//有缓冲通道

func request(url string) []byte {
	trans := http.Transport{TLSClientConfig: &tls.Config{InsecureSkipVerify: true}}
	client := http.Client{Transport: &trans}
	resp, err := client.Get(url)
	if err != nil {
		log.Println(err)
		return nil
	}
	defer resp.Body.Close()
	contents, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Println(err)
		return nil
	}
	return contents
}
func mirrorQuery() []byte {
	//response := make(chan []byte) //如果使用无缓冲通道，可能会出现goroutine泄露，这是一个bug，
	// 泄露的goroutine是不会被自动回收的
	//
	// 原因是，以下三个url由于访问时间关系，肯定
	//会有两个慢的，而这两个慢的就会被卡住，因为这个时候已经没有goroutine来接收他们的值了，已经reture了。

	response := make(chan []byte, 3)

	go func() {
		response <- request("https://www.baidu.com")
	}()

	go func() {
		response <- request("https://www.github.com")
	}()

	go func() {
		response <- request("https://www.360.net")
	}()
	return <-response
}

func getStack() {
	for tries := 0; time.Now().Before(time.Now().Add(time.Second * 10)); tries++ {
		fmt.Fprintf(os.Stderr, "%d\n", runtime.NumGoroutine())
		time.Sleep(time.Second * 1) //等一会，查看协程数量的变化
		fmt.Fprintf(os.Stderr, "%d\n", runtime.NumGoroutine())
	}
}

func getStackTraceHandler(w http.ResponseWriter, r *http.Request) {
	stack := debug.Stack()
	w.Write(stack)
	pprof.Lookup("goroutine").WriteTo(w, 2)
}

func main() {
	go getStack()
	fmt.Println(string(mirrorQuery()))
	time.Sleep(10 * time.Second)
}
