package main

import (
	"fmt"
	"links"
	"log"
)

//该版本的爬虫使用有缓存的通道作为计数信号量限制
//爬虫的并发数量

//令牌是一个计数信号量，确保并发请求限制在20个以内
var token = make(chan struct{}, 20) //当然这里也可以使用bool型或者int

func main() {
	worklist := make(chan []string)
	var n int //等待发送到任务列表的数量
	//从命令行开始
	n++
	//go func() {worklist<- os.Args[1:]}()
	go func() {
		worklist <- []string{"https://www.csdn.net"}
	}()

	seen := make(map[string]bool)
	for ; n > 0; n-- {
		list := <-worklist
		for _, link := range list {
			if !seen[link] {
				seen[link] = true
				n++
				go func(link string) {
					worklist <- crawl(link)
				}(link)
			}
		}
	}
}

func crawl(url string) []string {
	fmt.Println(url)
	token <- struct{}{} // 获取令牌,如果满了这里就会阻塞
	list, err := links.Extract(url)
	<-token //完成就从管道中取出一个，这样就会有新的协程抓取url了
	if err != nil {
		log.Print(err)
	}
	return list

}
