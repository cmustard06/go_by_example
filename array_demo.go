package main

import "fmt"

func main(){
	//声明一个长度为5 的整形数组
	var a[5] int
	fmt.Println("emp:",a)  // 默认为0

	a[4] = 100
	fmt.Println("set:", a)
	fmt.Println("get:", a[4])

	//获取长度
	fmt.Println("len: ", len(a))

	//声明一个数组，并初始化
	var b = [5]int{1,2,3,4,5}
	fmt.Println("decl:",b)

	//申明一个二维数组
	var twoD [2][3]int
	for i:=0;i<2;i++{
		for j:=0;j<3;j++{
			twoD[i][j] = i+j
		}
	}
	fmt.Println("2d:", twoD)
}
