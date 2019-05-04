package main

import (
	"bytes"
	"io"
)

var debug = true

func f(out io.Writer) {
	if out != nil { //如果声明的类型为*bytes.Buffer，那么这个表达式一直为true，但是如歌debug为false，那么其值为nil，调用Write方法就会宕机
		out.Write([]byte("done\n"))
	}
}

func main() {
	// var buf *bytes.Buffer   //这个会报异常的原因是因为这里的动态类型是*bytes.Buffer，因此在与nil做比较的时候一直都是不为nil
	var buf io.Writer //这里用接口的原因是因为没有初始化的情况下，它的动态类型和对应的值都为nil
	if debug {
		buf = new(bytes.Buffer)
	}
	f(buf)

}
