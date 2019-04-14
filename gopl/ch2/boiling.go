package main

import "fmt"

/*输出水的沸点*/

const boilingF = 212.0

func main() {
	var f = boilingF
	var c = (f - 32) * 5 / 9
	fmt.Printf("boling point = %gF or %g C\n", f, c) //浮点数输出，有效位数更多

}
