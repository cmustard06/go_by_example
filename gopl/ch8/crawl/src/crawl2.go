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
	var n int //等待发送到任务列表的数量，使用n跟踪任务数量，一个新的url就递增，爬取成功后就递减，这时候主循环会从n减到0，这时候就可以退出了
	//从命令行开始
	n++
	//go func() {worklist<- os.Args[1:]}()
	go func() {
		worklist <- []string{"https://www.csdn.net"}
	}()

	seen := make(map[string]bool)
	for ; n > 0; n-- { //如果不这样使用，主循坏将永远不会退出，因为如果没有新的url就会一直阻塞
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
