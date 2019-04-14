package main

import (
	"fmt"
	"log"
	"net/http"
)

/*迷你的echo服务器*/

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "URL.PATH=%q\n", r.URL.Path) //%q为输入带引号的字符串
	return
}
func main() {
	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe("127.0.0.1:8000", nil)) //日志出现错误，输出错误，然后结束程序

}
