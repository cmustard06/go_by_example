package main

import (
	"fmt"
	"net/http"
	"log"
)

//注册/list和/price
//!+handlerfunc
/*
package http
type HandlerFunc func(w ResponseWriter, r *Request)
func (f HandlerFunc) ServeHTTP(w ResponseWriter, r *Request) {
	f(w, r)
}
//!-handlerfunc
*/
type dollars float32

func (d dollars)String() string{
	return fmt.Sprintf("$%.2f",d)
}

type database map[string]dollars

func (db database)list(w http.ResponseWriter,r *http.Request){
	for item,price := range db{
		fmt.Fprintf(w,"%s:%s\n",item,price)
	}
}

func (db database) price(w http.ResponseWriter,r *http.Request){
	item := r.URL.Query().Get("item")
	price,ok := db[item]
	if !ok{
		//map中不存在
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprintf(w,"no such item%q\n",item)
		return
	}
	fmt.Fprintf(w,"%s\n",price)
	return
}

func main(){
	db := database{"shoes":50,"socks":5}
	mux := http.NewServeMux()
	mux.HandleFunc("/list",db.list)
	mux.HandleFunc("/price",db.price)
	log.Fatal(http.ListenAndServe("localhost:8000",mux))
}