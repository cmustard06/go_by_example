package main

import (
	"html/template"
	"os"
	"log"
)

/* 使用html/template 转义不受信任的string类型字符串
*/

func main(){
	const templ = `<p>A:{{.A}}</p><p>B:{{.B}}</p>`
	t := template.Must(template.New(templ).Parse(templ))
	var data struct{
		A string
		B template.HTML
	}
	data.A = "<b>Hello!</b>" //<p>A:&lt;b&gt;Hello!&lt;/b&gt;</p>  txt string 会被转义
	data.B = "<b>Hello!</b>"
	if err:= t.Execute(os.Stdout,data);err!=nil{
		log.Fatal(err)
	}
}
