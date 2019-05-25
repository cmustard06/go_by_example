package main

import (
	"flag"
	"os"
	"io/ioutil"
	"fmt"
	"path/filepath"
)

//创建一个程序来生成指定目录的硬盘使用情况报告，这个程序和Unix里的du工 具比较相似。
func main(){
	flag.Parse()
	roots := flag.Args()
	//如果没有，就使用默认的值
	if len(roots)==0{
		roots = []string{"."}
	}

	//遍历文件
	fileSizes := make(chan int64)
	go func() {
		for _,root := range roots{
			walkDir(root, fileSizes)
		}
		//关闭通道
		close(fileSizes)
	}()

	//打印结果
	var nfiles, nbytes int64
	for size := range fileSizes{
		nfiles ++
		nbytes +=size
	}
	fmt.Printf("%d files %.1f GB\n",nfiles,float64(nbytes)/1e9)
}
//以递归方式遍历以dir为根的文件树 ,并在fileSizes上发送每个找到的文件的大小。
func walkDir(dir string, fileSizes chan<- int64){
	for _,entry := range dirents(dir){
		if entry.IsDir(){
			//一个新的目录，继续深入
			subdir := filepath.Join(dir,entry.Name())
			walkDir(subdir,fileSizes)
		}else{
			//写入管道
			fileSizes<- entry.Size()
		}
	}
}

//返回目标目录下的entry
func dirents(dir string) []os.FileInfo{
	entries, err := ioutil.ReadDir(dir)
	if err!=nil{
		fmt.Fprintf(os.Stderr, "du1:%v\n",err)
		return nil
	}
	return entries
}