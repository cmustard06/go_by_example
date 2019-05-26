package main

import (
	"fmt"
	"os"
	"path/filepath"
	"sync"
	"time"
)

//增加停止的功能

var done = make(chan struct{})

func cacelled() bool {
	select {
	case <-done:
		return true
	default:
		return false
	}
}

func main() {
	roots := os.Args[1:]
	if len(roots) == 0 {
		roots = []string{"."}
	}

	//起一个协程实时监测输入，
	go func() {
		os.Stdin.Read(make([]byte, 1))
		close(done)
	}()

	fileSizes := make(chan int64)
	var n sync.WaitGroup
	for _, root := range roots {
		n.Add(1)
		go walkDir(root, &n, fileSizes)
	}

	go func() {
		n.Wait()
		close(fileSizes)
	}()

	tick := time.Tick(500 * time.Millisecond)
	var nfiles, nbytes int64
loop:
	for {
		select {
		case <-done:
			for range fileSizes {
				//就是为了取出通道中值
			}
			return
		case size, ok := <-fileSizes:
			if !ok {
				break loop
			}
			nfiles++
			nbytes += size
		case <-tick:
			fmt.Printf("%d files %.1f GB\n", nfiles, float64(nbytes)/1e9)

		}
	}
	fmt.Printf("%d files %.1f GB\n", nfiles, float64(nbytes)/1e9)

}

func walkDir(dir string, n *sync.WaitGroup, fileSizes chan<- int64) {
	defer n.Done()
	if cacelled() {
		return
	}

	for _, entry := range dirents(dir) {
		if entry.IsDir() {
			n.Add(1)
			subdir := filepath.Join(dir, entry.Name())
			go walkDir(subdir, n, fileSizes)
		} else {
			fileSizes <- entry.Size()
		}
	}

}

//控制协程数量
var sema = make(chan struct{}, 20)

func dirents(dir string) []os.FileInfo {
	select {
	case sema <- struct{}{}:
	case <-done: //这里判读是不人工停止了
		return nil
	}
	defer func() { <-sema }()

	file, err := os.Open(dir)
	if err != nil {
		fmt.Fprintf(os.Stderr, "du:%v\n", err)
		return nil
	}

	defer file.Close()
	infos, err := file.Readdir(0) //0代表没有限制，读取所有字段
	if err != nil {
		fmt.Fprintf(os.Stderr, "du: %v\n", err)
	}
	return infos

}
