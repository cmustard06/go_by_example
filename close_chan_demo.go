package main

import "fmt"

func main(){
	jobs := make(chan int, 5)
	done := make(chan bool)

	go func() {
		for{
			j, more := <-jobs
			if more{
				fmt.Println("recv job", j)
			}else{
				fmt.Println("recv all job")
				done<-true
				return
			}
		}
	}()

	for j:=1;j<=3;j++{
		jobs <- j
		fmt.Println("sent job", j)
	}
	close(jobs)
	fmt.Println("send all jobs")
	fmt.Printf("%+v",<- done)

}
