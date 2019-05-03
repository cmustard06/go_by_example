package main

import (
	"flag"
	"fmt"
	"time"
)

/*
./sleep -period 3s
*/
func main() {
	duration := flag.Duration("period", 1*time.Second, "sleep period") //标准包中自带的解析
	flag.Parse()
	fmt.Printf("Sleeping for %v...", *duration)
	time.Sleep(*duration)
	fmt.Println()
}
