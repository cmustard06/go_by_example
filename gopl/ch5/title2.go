package main

import (
	"golang.org/x/net/html"
	"fmt"
	"net/http"
	"strings"
)

//在title1的基础上，添加recover函数

func forEachNode(n *html.Node,pre,post func(n *html.Node)){
	if pre!=nil{
		pre(n)
	}
	for c:=n.FirstChild;c!=nil;c=c.NextSibling{
		forEachNode(c,pre,post)
	}
	if post!=nil{
		post(n)
	}
}

func soleTitle(doc *html.Node)(title string,err error){
	type bailout struct {}
	//recover，防止程序崩溃
	defer func() {
		switch p:=recover();p {
		case nil:
			//no panic
		case bailout{}:
			err = fmt.Errorf("多个title标签")
		default:
			panic(p)  //预料之外的panic
		}
	}()

	forEachNode(doc, func(n *html.Node) {
		if n.Type==html.ElementNode && n.Data=="title" && n.NextSibling!=nil{
			if title != ""{
				panic(bailout{}) // 如果存在多个title标签就抛异常
			}
			title = n.FirstChild.Data
		}
	},nil)
	if title == ""{
		return "",fmt.Errorf("no title Element")
	}
	return title,nil
}


func title(url string) error {
	resp, err := http.Get(url)
	if err != nil {
		return err
	}

	// Check Content-Type is HTML (e.g., "text/html; charset=utf-8").
	ct := resp.Header.Get("Content-Type")
	if ct != "text/html" && !strings.HasPrefix(ct, "text/html;") {
		resp.Body.Close()
		return fmt.Errorf("%s has type %s, not text/html", url, ct)
	}

	doc, err := html.Parse(resp.Body)
	resp.Body.Close()
	if err != nil {
		return fmt.Errorf("parsing %s as HTML: %v", url, err)
	}
	title, err := soleTitle(doc)
	if err != nil {
		return err
	}
	fmt.Println(title)
	return nil
}

func main() {
	title("http://192.168.1.197/dest/login.html")
}