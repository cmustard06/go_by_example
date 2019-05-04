package main

import (
	"fmt"
	"log"
	"net/http"
)

//通过http.Handler接口中的serveHTTP函数实现一个自定义路由
type dollars float32

func (d dollars) String() string {
	return fmt.Sprintf("$%.2f", d)
}

type database map[string]dollars

func (db database) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	switch r.URL.Path {
	case "/list":
		for item, price := range db {
			fmt.Fprintf(w, "%s:%s\n", item, price)
		}
	case "/price":
		item := r.URL.Query().Get("item")
		price, ok := db[item]
		if !ok {
			w.WriteHeader(404)
			fmt.Fprintf(w, "no such item:%q\n", item) //输出带引号的字符串
			return
		}
		fmt.Fprintf(w, "%s\n", price)
	default:
		w.WriteHeader(404)
		fmt.Fprintf(w, "no such page:%s\n", r.URL)
	}
}

func main() {
	db := database{"shoes": 50, "socks": 5}
	log.Fatal(http.ListenAndServe("localhost:8000", db))
}
