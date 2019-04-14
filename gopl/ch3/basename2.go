package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

/* 改进basename1中获取basname的方式，采用strings中
的函数截断
*/

func basename(s string) string {
	slash := strings.LastIndex(s, "/")
	s = s[slash+1:]
	if dot := strings.LastIndex(s, "."); dot >= 0 {
		s = s[:dot]
	}
	return s
}

func main() {
	input := bufio.NewScanner(os.Stdin)
	for input.Scan() {
		if input.Text() == "exit" {
			break
		}

		fmt.Println(basename(input.Text()))
	}
}
