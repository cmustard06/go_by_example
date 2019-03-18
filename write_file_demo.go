package main

import (
	"io/ioutil"
	"os"
	"fmt"
	"bufio"
)

func check(err error) {
	if err != nil {
		panic(err)
	}
}

func main() {
	d1 := []byte("hello\ngo\n")
	err := ioutil.WriteFile("write_demo.txt", d1, 0644)
	check(err)
	//使用较底层的方法
	f, err := os.Create("write_demo.txt")
	check(err)

	defer f.Close() //在该函数结束之前调用该函数
	// 追加内容
	d2 := []byte{115, 111, 109, 101, 10}
	n2, err := f.Write(d2)
	check(err)
	fmt.Printf("wrote %d bytes\n", n2)
	//追加
	n3, err := f.WriteString("write\n")
	fmt.Printf("wrote %d bytes\n", n3)
	//刷新writer，将内存中新写入的数据刷新到硬盘中
	f.Sync()
	w := bufio.NewWriter(f)
	n4, err := w.WriteString("bufferd123\n")
	fmt.Printf("wrote %d bytes\n", n4)
	//确保所有缓冲操作都已成功应用
	w.Flush()

}
