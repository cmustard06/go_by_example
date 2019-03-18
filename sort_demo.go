package main

import (
	"sort"
	"fmt"
)

func main(){
	str := []string{"c","a","b"}
	sort.Strings(str)
	fmt.Println("strings:",str)

	ints := []int{7,2,4}
	sort.Ints(ints)
	fmt.Println("ints:",ints)

	s := sort.IntsAreSorted([]int{3,1,2})  //测试是否排序完成
	fmt.Println("sorted:",s)
}
