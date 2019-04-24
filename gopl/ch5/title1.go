package main

import (
	"golang.org/x/net/html"
	"net/http"
	"crypto/tls"
	"fmt"
)

//指定一个url打印其title

func forEachNode(n *html.Node,pre,post func(n *html.Node)){
	if pre!=nil{
		pre(n)
	}
	for c:= n.FirstChild;c!=nil;c=c.NextSibling{
		forEachNode(c,pre,post)
	}
	if post!=nil{
		post(n)
	}
}

func title(url string) error{
	tr := http.Transport{TLSClientConfig:&tls.Config{InsecureSkipVerify:false}}
	client := http.Client{Transport:&tr}
	resp, err := client.Get(url)
	if err!=nil{
		return err
	}
	defer resp.Body.Close()
	node, err := html.Parse(resp.Body)
	if err!=nil{
		return nil
	}
	visitNode := func(n *html.Node) {
		if n.Type == html.ElementNode && n.Data=="title" && n.FirstChild!=nil{
			fmt.Println(n.FirstChild.Data)
		}
	}
	//forEachNode(node,visitNode,nil)
	forEachNode(node,nil,visitNode)
	return nil
}

func main(){
	url := "https://www.csdn.net"
	title(url)
}