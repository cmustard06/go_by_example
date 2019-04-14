package main

import (
	"bufio"
	"fmt"
	"os"
)

/*从stdin读取文件名称，然后打印每一个文件的basename
 */

//eg. a=>a,a.go->a a/b/c.go->c

func basename(s string) string {
	for i := len(s) - 1; i >= 0; i-- {
		if s[i] == '/' {
			s = s[i+1:]
			break
		}
	}

	for i := len(s) - 1; i >= 0; i-- {
		if s[i] == '.' {
			s = s[:i]
			break
		}
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
