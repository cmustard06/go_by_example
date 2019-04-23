package main

import (
	"golang.org/x/net/html"
	"os"
	"fmt"
	"net/http"
	"crypto/tls"
	"time"
)

//从一个HTML文档中抓取href，并输出


func visit(links []string, n *html.Node)[]string{
	if n.Type == html.ElementNode && n.Data == "a"{
		for _,a := range n.Attr{
			if a.Key == "href"{
				//fmt.Println(a.Val)
				links = append(links,a.Val)
			}
		}
	}
	for c := n.FirstChild;c!=nil;c=c.NextSibling{
		links = visit(links,c)
	}
	return links
}

func main(){
	tr := &http.Transport{TLSClientConfig:&tls.Config{InsecureSkipVerify:false}}
	client := &http.Client{Transport:tr, Timeout:time.Second*2}
	resp,err := client.Get("https://docs.python.org/3.7/library/index.html")
	if err!=nil{
		fmt.Errorf("get error %v",err)
	}
	defer resp.Body.Close()

	doc,err := html.Parse(resp.Body)
	if err!=nil{
		fmt.Fprintf(os.Stderr,"findlinks1:%v\n",err)
		os.Exit(1)
	}
	for _,link := range visit(nil,doc){
		fmt.Println(link)
	}
}