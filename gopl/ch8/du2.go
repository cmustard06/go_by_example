package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"time"
)

//在版本1的基础上增加实时进度动态

var verbose = flag.Bool("v", false, "展示详细进度信息")

func main() {
	//启动后台进程，定期输出结果
	flag.Parse()
	roots := flag.Args()
	if len(roots) == 0 {
		roots = []string{"."}
	}

	//遍历文件
	fileSizes := make(chan int64)
	go func() {
		for _, root := range roots {
			walkDir(root, fileSizes)
		}
		close(fileSizes)
	}()

	//定期打印结果
	var tick <-chan time.Time

	if *verbose {
		tick = time.Tick(500 * time.Millisecond)
	}
	var nfiles, nbytes int64

loop:
	for {
		select {
		case size, ok := <-fileSizes: //这里没有使用range，所以只能进行显示判断管道是否关闭
			if !ok {
				break loop //管道已经被关闭，直接跳出for循环，没有loop只能跳出select
			}
			nfiles++
			nbytes += size
		case <-tick:
			fmt.Printf("%d files %.1f GB\n", nfiles, float64(nbytes)/1e9)
		}

	}
	//最后
	fmt.Printf("%d files %.1f GB\n", nfiles, float64(nbytes)/1e9)
}

func walkDir(dir string, fileSizes chan<- int64) {
	for _, entry := range dirents(dir) {
		if entry.IsDir() {
			sudir := filepath.Join(dir, entry.Name())
			walkDir(sudir, fileSizes)
		} else {
			fileSizes <- entry.Size()
		}
	}
}

func dirents(dir string) []os.FileInfo {
	entries, err := ioutil.ReadDir(dir)
	if err != nil {
		fmt.Fprintf(os.Stderr, "du: %v\n", err)
		return nil
	}
	return entries
}
