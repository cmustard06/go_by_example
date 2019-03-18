package main

import "fmt"

func main(){
	//创建一个slice，对比数组，切片是需要固定长度的
	var s = make([]string, 3)
	//fmt.Printf("%T",s)
	//s = append(s, "qq","mmm")
	fmt.Println(s)
	s[0] = "a"
	s[1] = "n"
	s[2] = "c"
	fmt.Println("set:", s)
	fmt.Println("get: ", s[2])

	s = append(s, "d")
	s = append(s, "e", "f")
	fmt.Println("apd:", s)

	c := make([]string, len(s))
	copy(c, s)
	fmt.Println("cpy: ",c)

	l := s[2:5]
	fmt.Println("sl1:", l)
}
