package main

import (
	"net/http"
	"io/ioutil"
	"fmt"
	"bytes"
	"net/url"
	"strings"
	"time"
	//"io"
	//"bufio"
	"log"
)

func http_get(){
	resp, err := http.Get("http://192.168.1.197")
	//异常处理
	defer func(){
		resp.Body.Close()
		e := recover()
		if e==nil{
			return
		}else{
			fmt.Println(e)
		}
	}()
	if err!=nil{
		fmt.Println(err)
	}
	data,err := ioutil.ReadAll(resp.Body)
	if err!=nil{
		fmt.Println(err)
	}
	fmt.Println(string(data))

}

func http_post(){
	buf := bytes.NewBuffer([]byte("a=1&b=2"))

	resp, _ := http.Post("http://192.168.13.180", "application/x-www-form-urlencoded", buf)
	defer func() {
		resp.Body.Close()

		e := recover()
		if e==nil{
			return
		}else{
			fmt.Println(e)
		}
	}()
	data, _ := ioutil.ReadAll(resp.Body)
	fmt.Println(string(data))


}


func http_cus(){
	v := url.Values{}
	v.Set("username","xxxx")
	v.Set("password","xxxxx")
	//将form 数据编码
	body:= ioutil.NopCloser(strings.NewReader(v.Encode()))
	client := &http.Client{
		Timeout:time.Second,
	} //客户端，被Get，Head以及Post使用
	request, e := http.NewRequest("POST", "http://192.168.13.180", body)
	if e!=nil{
		fmt.Println(e)
	}

	//设置http请求头部
	request.Header.Set("Content-Type", "application/x-www-form-urlencoded;param=value") //设置完成该头部后，post参数才能正常提交

	resp, err := client.Do(request)
	if err!=nil{
		panic(err)
	}
	defer resp.Body.Close()
	content,err := ioutil.ReadAll(resp.Body)
	fmt.Println(resp.Header.Get("Date"))
	if err!=nil{
		panic(err)
	}
	fmt.Println(string(content))

}

func hijk(){
	http.HandleFunc("/hijack", func(w http.ResponseWriter, r *http.Request) {
		hj,ok := w.(http.Hijacker)
		if !ok{
			http.Error(w, "webserver doesn't support hijacking",http.StatusInternalServerError)
			return
		}
		conn, writer, e := hj.Hijack()
		if e!=nil{
			http.Error(w, e.Error(),http.StatusInternalServerError)
			return
		}
		defer conn.Close()
		writer.WriteString("Now we're speaking raw TCP")
		writer.Flush()

		s, err := writer.ReadString('\n')
		if err!=nil{
			log.Printf("error reading string %v",err)
			return
		}
		fmt.Fprintf(writer,"HTTP/1.1 200 OK\r\n\r\nyou said:%q\nBye.\n",s)
		writer.Flush()

	})
}


func main(){
	//http_get()
	//http_post()
	http_cus()
	//hijk()
	//http.ListenAndServe(":8080",nil)
}