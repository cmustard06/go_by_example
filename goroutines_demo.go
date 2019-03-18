package main

import "fmt"

func f(from string){
	for i:=0;i<10;i++{
		fmt.Println(from,":",i)
	}
}

func main(){
	go f("direct")
	go f("goroutine")

	go func(msg string) {
		fmt.Println(msg)
	}("going")

	fmt.Scanln()
	fmt.Println("done")
}
