package main

import (
	"net/http"
	"crypto/tls"
	"golang.org/x/net/html"
	"fmt"
)
//获取目标网站的标签树状图
var depth = 2

func outline2(url string) error{
	tr := http.Transport{TLSClientConfig:&tls.Config{InsecureSkipVerify:false}}
	client := http.Client{Transport:&tr}
	resp, err := client.Get(url)
	if err!=nil{
		return err
	}
	defer resp.Body.Close()
	node, err := html.Parse(resp.Body)
	if err!=nil{
		return err
	}
	forEachNode(node,startElement,endElement)
	return nil
}

func forEachNode(n *html.Node, pre,post func(n *html.Node)){
	if pre!=nil{
		pre(n)   //开始标签如<html>
	}
	for c:=n.FirstChild;c!=nil;c=c.NextSibling{
		forEachNode(c,pre,post)
	}
	if post!=nil{
		post(n)  //结束标签如</html>
	}
}

func startElement(n *html.Node){
	if n.Type==html.ElementNode{
		fmt.Printf("%*s<%s>\n",depth*2,"",n.Data) //缩进长度
		depth++
	}
}

func endElement(n *html.Node){
	if n.Type == html.ElementNode{
		depth--
		fmt.Printf("%*s</%s>\n",depth*2,"",n.Data)
	}
}

func main(){
	urls := []string{"https://www.aisec.com","https:www.hao123.com"}
	for _,url := range urls{
		outline2(url)
	}
}