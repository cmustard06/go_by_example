package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"sync"
	"time"
)

//一个文件夹起一个goroutine

var verbose = flag.Bool("v", false, "展示详细详细")

func main() {
	flag.Parse()

	roots := flag.Args()
	if len(roots) == 0 {
		roots = []string{"."}
	}

	//并行运行
	fileSizes := make(chan int64)
	var n sync.WaitGroup //对协程进行监控，记录协程实例
	for _, root := range roots {
		//增加一个协程
		n.Add(1)
		go walkDir(root, &n, fileSizes)
	}
	go func() {
		n.Wait() //等待所以协程完成之后，关闭通道
		fmt.Println("等待完成")
		close(fileSizes)
	}()

	var tick <-chan time.Time
	if *verbose {
		tick = time.Tick(500 * time.Millisecond)
	}
	var nfiles, nbytes int64

loop:
	for {
		select {
		case size, ok := <-fileSizes:
			if !ok {
				break loop // 通道已经关闭
			}
			nfiles++
			nbytes += size
		case <-tick:
			fmt.Printf("%d files %.1f GB\n", nfiles, float64(nbytes)/1e9)
		}
	}
	fmt.Printf("%d files %.1f GB\n", nfiles, float64(nbytes)/1e9)
}

func walkDir(dir string, n *sync.WaitGroup, fileSize chan<- int64) {
	defer n.Done() //完成一个，减少一个
	for _, entry := range dirents(dir) {
		if entry.IsDir() {
			//如果是，再起一个
			n.Add(1)
			subdir := filepath.Join(dir, entry.Name())
			fmt.Println(subdir)
			go walkDir(subdir, n, fileSize)
		} else {
			fmt.Println(filepath.Join(dir, entry.Name()))
			fileSize <- entry.Size()
		}
	}

}

//由于在程序高峰时可能会出现上千个goroutine，因此
//我们不得不使用计数器信号量做限制
var sema = make(chan struct{}, 20)

func dirents(dir string) []os.FileInfo {
	sema <- struct{}{} //如果满了就会阻塞
	defer func() {
		<-sema //函数跑完了，就释放
	}()
	infos, err := ioutil.ReadDir(dir)
	if err != nil {
		fmt.Fprintf(os.Stderr, "du: %n", err)
		return nil
	}
	return infos
}
