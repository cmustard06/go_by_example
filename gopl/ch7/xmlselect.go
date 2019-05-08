package main

import (
	"encoding/xml"
	"os"
	"io"
	"fmt"
	"strings"
)

//打印XML文档的选定元素的文本

func main(){
	dec := xml.NewDecoder(os.Stdin)
	var stack []string //元素名称的栈
	for{
		token, err := dec.Token()
		if err==io.EOF{
			break
		}else if err!=nil{
			fmt.Fprintf(os.Stderr,"xmlselect: %v\n",err)
			os.Exit(1)
		}
		switch tok := token.(type) {
		case xml.StartElement:
			stack =  append(stack,tok.Name.Local)
		case xml.EndElement:
			stack = stack[:len(stack)-1] //pop
		case xml.CharData:
			if containsAll(stack,os.Args[1:]){
				fmt.Printf("%s:%s\n",strings.Join(stack," "),tok)

			}
		}
	}
}

func containsAll(x,y []string) bool {
	for len(y) <= len(x){
		if len(y) == 0{
			return true
		}
		if x[0]==y[0]{
			y = y[1:]
		}
		x = x[1:]
	}
	return false
}