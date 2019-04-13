package main

import (
	"bufio"
	"fmt"
	"os"
)

/*
输出标准输入中出现次数大于1的行，前面是次数,输入exit即可退出
*/

func main() {
	counts := make(map[string]int)
	input := bufio.NewScanner(os.Stdin)
	for input.Scan() {
		if input.Text() == "exit" {
			break
		}
		counts[input.Text()]++
	}

	for line, n := range counts {
		if n > 1 {
			fmt.Printf("%d\t%s\n", n, line)
		}
	}
}
