package main

import (
	"io/ioutil"
	"fmt"
	"os"
	"io"
	"bufio"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	//ioutil更方便的读取文件内容
	dat, err := ioutil.ReadFile("test.go")
	check(err)
	fmt.Println(string(dat))

	f, err := os.Open("test.go")
	check(err)

	b1 := make([]byte, 5)
	n1, err := f.Read(b1)
	check(err)
	fmt.Printf("%d bytes: %s\n", n1, string(b1))

	o2, err := f.Seek(6, 0)
	check(err)
	b2 := make([]byte, 2)
	n2, err := f.Read(b2)
	check(err)
	fmt.Printf("%d bytes @ %d: %s\n", n2, o2, string(b2))

	o3, err := f.Seek(6, 0)
	check(err)
	b3 := make([]byte, 2)
	//使用io包可以更方便实现读取文件的操作
	n3, err := io.ReadAtLeast(f, b3, 2)
	check(err)
	fmt.Printf("%d bytes @ %d: %s\n", n3, o3, string(b3))

	_, err = f.Seek(0, 0)
	check(err)
	//bufio包实现了一个缓冲读取器
	r4 := bufio.NewReader(f)
	//指定读取的字节数
	b4,err := r4.Peek(5)
	check(err)
	fmt.Printf("5 bytes:%s\n",string(b4))

	f.Close()

}
