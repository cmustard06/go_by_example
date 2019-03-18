package main

import "fmt"

//计算两个数的最大公约数
func fib(n int) int {
	x, y := 0, 1
	for i := 0; i < n; i++ {
		x, y = y, x+y
	}
	return x
}

var global *int

func f() {
	var x int
	x = 1
	global = &x // 将x从栈上移动到堆上
}

func g() {
	var y *int //在栈上分配空间
	*y = 1
}

func main() {
	var 无0 = 0 //只要是Unicode就可以使用
	fmt.Println(无0)

	// 结构体、数组、slice中的元素时可寻址的
	names := []string{"ken", "robe"}
	fmt.Printf("%v\n", &names[1]) //0xc00004a430

	type User struct {
		Name string
		Age  int
	}

	ken := User{"ken", 18}
	fmt.Printf("%v\n", &ken.Age)

	//map是不可寻址的，元素可能会随着map的增长而动态改变地址

	dic := make(map[string]string)
	dic["one"] = "--"
	//fmt.Printf("%v\n", &dic["one"]) //cannot take the address of dic["one"]

	// 编译器从最内层开始寻找变量的声明
	// 外层的同名声明会被覆盖
	s := "hello!" // string
	for i := 0; i < len(s); i++ {
		s := s[i]
		if s != '!' {
			s := s + 'A' - 'a'
			fmt.Printf("%c", s)
		}
	}
	fmt.Println(s)

	//switch 同if，可带初始化语句
	switch x := 0; x > -1 {
	case true:
		sn := 1
		fmt.Println(sn)
	default:
		// fmt.Println(sn)  //各case之间相互独立
	}

}
