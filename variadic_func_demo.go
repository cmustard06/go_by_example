package main

import "fmt"

func sum(nums ...int){
	fmt.Println(nums," ")
	total := 0
	for _,num := range nums{
		total += num
	}
	fmt.Println(total)
}

func main(){
	sum(1,2)
	sum(1,2,3)

	num := []int{1,2,3,4}
	sum(num...)  // 可以直接传入数组，但是注意使用...语法做标注

}
