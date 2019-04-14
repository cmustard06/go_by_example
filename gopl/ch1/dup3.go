package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

/*简化程序，这里只读取文件的内容，不在读取标准输入的内容，读取文件不在使用open函数，
改为更高级的函数*/

func main() {
	counts := make(map[string]int)
	for _, filename := range os.Args[1:] {
		data, err := ioutil.ReadFile(filename) //更高级的文件内容读取函数
		if err != nil {
			fmt.Fprintf(os.Stderr, "dup3: %v\n", err)
			continue
		}
		for _, line := range strings.Split(string(data), "\n") {
			counts[line]++
		}
	}
	for line, n := range counts {
		if n > 1 {
			fmt.Printf("%d\t%#v\n", n, line)
		}
	}
}
