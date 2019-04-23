package main

import (
	"crypto/tls"
	"fmt"
	"golang.org/x/net/html"
	"log"
	"net/http"
	"net/url"
	"time"
)

//简单的网页url爬虫，待完善

func visitNode(links []string, n *html.Node) []string {
	if n.Type == html.ElementNode && n.Data == "a" {
		for _, a := range n.Attr {
			if a.Key != "href" {
				continue
			}
			link, err := url.Parse(a.Val)
			if err != nil {
				continue
			}
			links = append(links, link.String())
		}
	}
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		links = visitNode(links, c)
	}
	return links
}

func Extract(url string) ([]string, error) {
	tr := http.Transport{TLSClientConfig: &tls.Config{InsecureSkipVerify: false}}
	client := http.Client{Timeout: time.Second * 4, Transport: &tr}
	resp, err := client.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("getting %s:%s", url, resp.Status)
	}
	node, err := html.Parse(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("parsing %s as HTML:%v", url, err)
	}

	links := visitNode(nil, node)
	return links, nil
}

func crawl(url string) []string {
	fmt.Println(url) //输出爬到的url
	list, err := Extract(url)
	//fmt.Println(list)
	if err != nil {
		log.Print(err)
	}
	return list
}

func breadthFirst(f func(item string) []string, worklist []string) {
	seen := make(map[string]bool)
	for len(worklist) > 0 {
		items := worklist
		worklist = nil //注意作用域
		for _, item := range items {
			if !seen[item] {
				seen[item] = true //表示该页已经被爬过了
				//这里还可以对爬虫的深度进行设置
				worklist = append(worklist, f(item)...) //将爬到的新链接都保存到表中
			}
		}
	}
}

func main() {
	breadthFirst(crawl, []string{"https://www.csdn.net", "https://demo.aisec.cn"})
}
