package main

import (
	"fmt"
	"strconv"
	// "unicode/utf8"
)

func main() {
	//byte 是uint8 的别名，强调值是原始字节数据
	var b1 uint8
	var b2 byte

	fmt.Println(b1 == b2) //true

	// rune是int32的别名，强调的是Unicode码点值
	var r1 rune
	var i1 int32

	fmt.Println(i1 == r1) //true

	// int 与 int32 在绝大部分平台上是一致的，但编译器不认为是同一类型
	//

	var i2 uint8 = 127
	fmt.Println(strconv.FormatInt(127*127, 2)) //11111100000001 转化为二进制
	fmt.Println(i2 * i2)                       //1 // 发生溢出后，高于 8 位的直接截取掉，乘积剩下 0000 0001
	fmt.Println(100 >> 3)                      // 左移运算符结果向下取整12.5 ->12
}
