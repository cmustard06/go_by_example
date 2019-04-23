package main

import (
	"crypto/tls"
	"fmt"
	"golang.org/x/net/html"
	"net/http"
	"time"
)

func forEachNode(n *html.Node, pre, post func(n *html.Node)) {
	if pre != nil {
		pre(n)
	}
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		forEachNode(c, pre, post) //迭代查询
	}
	if post != nil {
		post(n)
	}
}

func Extract(url string) ([]string, error) {
	tr := http.Transport{TLSClientConfig: &tls.Config{InsecureSkipVerify: false}}
	client := &http.Client{Transport: &tr, Timeout: 4 * time.Second}
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
		return nil, fmt.Errorf("parsing %s as Html: %v", url, err)
	}

	var links []string
	visitNode := func(n *html.Node) {
		if n.Type == html.ElementNode && n.Data == "a" {
			for _, a := range n.Attr {
				if a.Key != "href" {
					continue
				}
				link, err := resp.Request.URL.Parse(a.Val)
				if err != nil {
					continue // 忽略bad urls
				}
				links = append(links, link.String())
			}
		}
	}
	forEachNode(node, visitNode, nil) //对于函数，默认为引用传递
	return links, nil
}

func main() {
	url := "https://www.python.org"
	strings, err := Extract(url)
	if err != nil {
		fmt.Errorf("error is %v", err)
	}
	fmt.Println(strings)
}
