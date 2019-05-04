package main

import (
	"fmt"
	"log"
	"net/http"
)

type dollars float32

func (d dollars) String() string {
	return fmt.Sprintf("$%.2f", d)
}

/*
type Handler interface {
	ServeHTTP(w ResponseWriter, r *Request)
}

*/
type database map[string]dollars

func (db database) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	for item, price := range db {
		fmt.Fprintf(w, "%s:%s\n", item, price)
	}
}

func main() {
	db := database{"shoe": 50, "socks": 5}
	log.Fatal(http.ListenAndServe("localhost:8000", db))
}
