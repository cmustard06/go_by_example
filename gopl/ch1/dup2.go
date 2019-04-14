package main

import (
	"bufio"
	"fmt"
	"os"
)

/*输出标准输入或者文件中出现次数大于1的行，前面是次数,输入exit即可退出*/

func countLines(f *os.File, counts map[string]int) {
	input := bufio.NewScanner(f)
	for input.Scan() {
		if input.Text() == "exit" {
			break
		}
		counts[input.Text()]++
	}

}
func main() {
	counts := make(map[string]int)
	files := os.Args[1:]
	if len(files) == 0 {
		countLines(os.Stdin, counts)
	} else {
		for _, arg := range files {
			f, err := os.Open(arg)
			defer f.Close()
			if err != nil {
				fmt.Fprintf(os.Stderr, "dup2 :%v\n", err)
				continue
			}
			countLines(f, counts)
		}
	}
	for line, n := range counts {
		if n > 1 {
			fmt.Printf("%d\t%#v\n", n, line) //#v 表示输出go语言格式的字符串
		}
	}
}
