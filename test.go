package main

import "fmt"

type Bird struct {
	Age int
	Name string
}

func passV(b *Bird){
	b.Age++
	b.Name = "Great"+b.Name
	fmt.Printf("传入修改后的Bird:\t\t %+v, \t\t 内存地址:%p,指针的内存地址: %p\n", *b, b, &b)
}

func main(){
	parrot := &Bird{Age:1, Name:"bule"}
	fmt.Printf("原始的Bird:\t\t %+v, \t\t 内存地址:%p,指针的内存地址: %p\n", *parrot, parrot, &parrot)
	passV(parrot)
	fmt.Printf("调用后Bird:\t\t %+v, \t\t 内存地址:%p,指针的内存地址: %p\n", *parrot, parrot, &parrot)
}
