package main

import (
	"fmt"
	"time"
)

func worker(id int, jobs<- chan int, result chan <-int){
	for j:= range jobs{
		fmt.Println("worker",id,"started job")
		time.Sleep(time.Second)
		fmt.Println("worker", id,"finish job", j)
		result <- j*2
	}
}

func main(){
	jobs := make(chan int,100)
	results := make(chan int, 100)

	for w:=1;w<=3;w++{
		go worker(w,jobs,results)
	}

	for j:=1;j<=5;j++{
		jobs<- j
	}
	//关闭jobs的管道
	close(jobs)

	for a:=1;a<=5;a++{
		<- results
		//fmt.Println("results:",<- results)
	}
}
