package main

import (
	"crypto/tls"
	"fmt"
	"golang.org/x/net/html"
	"log"
	"net/http"
	"time"
)

//html的简单实例

func outline(stack []string, n *html.Node) {
	if n.Type == html.ElementNode {
		stack = append(stack, n.Data)
		fmt.Println(stack) // [html body div div div ul li a] n.Data就是标签
	}

	for c := n.FirstChild; c != nil; c = c.NextSibling {
		outline(stack, c)
	}
}

func main() {
	tr := http.Transport{TLSClientConfig: &tls.Config{InsecureSkipVerify: false}}
	client := http.Client{Transport: &tr, Timeout: time.Second * 3}
	resp, err := client.Get("https://www.aisec.com")
	if err != nil {
		log.Print(err)
	}
	defer resp.Body.Close()
	node, err := html.Parse(resp.Body)
	if err != nil {
		log.Print(err)
	}
	outline(nil, node)
}
