package main

import (
	"bufio"
	"os"
	"fmt"
)

/*
打印输出的每一行，如果是重复的就不输出
*/

func main(){
	var a int64 = 24*60*60*1000*1000
	var b int64 = 24*60*60*1000
	fmt.Println(a/b)
	seen := make(map[string]bool)
	input := bufio.NewScanner(os.Stdin)
	for input.Scan(){
		line := input.Text()
        if !seen[line]{
        	seen[line] = true
        	fmt.Println(line)
		}
	}
	if err:= input.Err();err!=nil{
		fmt.Fprintf(os.Stderr,"dedup: %v \n",err)
		os.Exit(1)
	}
	fmt.Println(seen)
}