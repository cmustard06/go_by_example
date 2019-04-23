package main

import (
	"crypto/tls"
	"fmt"
	"golang.org/x/net/html"
	"net/http"
	"time"
)

//单个页面抓取链接

func visit(links []string, n *html.Node) []string {
	if n.Type == html.ElementNode && n.Data == "a" {
		for _, a := range n.Attr {
			if a.Key == "href" {
				links = append(links, a.Val)
			}
		}
	}
	//迭代查询href直到在本页中查询完成
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		links = visit(links, c)
		//fmt.Println(c.Attr)
	}
	return links
}

func main() {
	tr := http.Transport{TLSClientConfig: &tls.Config{InsecureSkipVerify: false}}
	client := http.Client{Transport: &tr, Timeout: time.Second * 4}
	resp, err := client.Get("https://www.csdn.net")
	if err != nil {
		fmt.Errorf("get error: %#v", err)
	}
	defer resp.Body.Close()
	node, err := html.Parse(resp.Body)
	if err != nil {
		fmt.Errorf("html parse err:%#v", err)
	}
	links := visit(nil, node)
	for _, link := range links {
		fmt.Println(link)
	}

}
