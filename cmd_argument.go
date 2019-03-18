package main

import (
	"os"
	"fmt"
)

func main(){
	argsWithProg := os.Args   //包括程序全路径名
	argwithoutProg := os.Args[1:] //不包括程序全路径名

	arg := os.Args[3]

	fmt.Println(argsWithProg)
	fmt.Println(argwithoutProg)
	fmt.Println(arg)
}
