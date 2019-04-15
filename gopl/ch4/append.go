package main

import "fmt"

/* 手动实现append函数

*/

//实现slice类型的append函数
func appendSlice(x []int, y ...int) []int {
	var z []int
	zlen := len(x)+len(y)
	if zlen <= cap(x){
		z = x[:zlen]
	}else{
		//空间不够，从新分配一个2倍的空间
		zcap := zlen
		if zcap < 2*len(x){
			zcap = 2*len(x)
		}

		z = make([]int,zlen,zcap)
		copy(z,x)
	}
	copy(z[len(x):],y)
	return z
}

//实现增加一个数据的append
func appendInt(x []int, y int)[]int{
	var z []int
	zlen := len(x)+1
	if zlen <= cap(x){
		z = x[:zlen]
	}else{
		zcap := zlen
		if zcap<2*len(x){
			zcap = 2*len(x)
		}
		z = make([]int,zlen,zcap)
		copy(z,x)

	}
	z[len(x)]=y
	return z
}

func main(){
	var x,y []int
	for i:=0;i<10;i++{
		y = appendInt(x,i)
		fmt.Printf("%d cap=%d\t%v\n",i,cap(y),y)
		x = y
	}
}