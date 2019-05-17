package main

import (
	"fmt"
	"links"
	"log"
	"os"
)

/*
由于并行程度太高了，这样会造成创建太多连接，超过了
程序能打开文件数的限制，导致dns查询失败或者连接失败
我们需要限制程序的并发数量
*/

func crawl(url string) []string {
	fmt.Println(url)
	list, err := links.Extract(url)
	if err != nil {
		log.Print(err)
	}
	return list

}

func main() {
	worklist := make(chan []string)
	//从命令行参数开始
	go func() { worklist <- os.Args[1:] }()

	//并发爬取Web
	seen := make(map[string]bool) //已经爬取的url
	for list := range worklist {
		for _, link := range list {
			if !seen[link] {
				seen[link] = true //现在爬取，先职位true
				go func(link string) {
					worklist <- crawl(link)
				}(link)
			}
		}
	}
}
