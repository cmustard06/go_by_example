package main

import (
	"fmt"
	"links"
	"log"
)

//版本2的替代方案，这里取消了计数器信号量，而是使用了一种替代方案（extract有待完善。什么链接都取）。

func crawl(url string) []string {
	fmt.Println(url)
	list, err := links.Extract(url)
	if err != nil {
		log.Print(err)
	}
	return list
}

func main() {
	worklist := make(chan []string)  //可能有重复的url列表
	unseenLinks := make(chan string) //去重后的url列表

	//向任务列表中添加命令行参数
	go func() {
		worklist <- []string{"http://gopl.io"}
	}()

	//创建20个协程来获取每一个不可见链接
	for i := 0; i < 20; i++ {
		go func() {
			for link := range unseenLinks {
				foundLinks := crawl(link)
				go func() { worklist <- foundLinks }() //这里阻塞不会影响20个协程的一个
				//worklist<-foundLinks //这样可能会一直阻塞一个goroutine，导致可运行的协程越来越少
			}
		}()
	}
	//主协程对URL进行去重操作，并把没有爬取的url发送给爬虫程序
	seen := make(map[string]bool)
	for list := range worklist {
		for _, link := range list {
			if !seen[link] {
				seen[link] = true
				unseenLinks <- link
			}
		}
	}
}
