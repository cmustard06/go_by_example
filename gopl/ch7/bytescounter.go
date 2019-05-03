package main

import (
	"bytes"
	"fmt"
)

//为类型 ByteCOunter 实现一个wirte方法用以统计bytes数量，

type ByteCounter int

/*
type Writer interface {
	Write(p []byte) (n int, err error)
}
*/
func (c *ByteCounter) Write(p []byte) (int, error) {
	*c += ByteCounter((len(p)))
	return len(p), nil
}

/*
type Stringer interface {
	String() string
}*/

type ByteString []byte

func (c ByteString) String() string {
	var buf bytes.Buffer
	buf.WriteByte('{')
	for _, v := range c {
		buf.WriteByte(v)
		buf.WriteByte(',')
	}
	buf.WriteByte('}')
	return buf.String()
}

func main() {
	var c ByteCounter
	c.Write([]byte("hello"))
	fmt.Println(c) //5 ， 该类型满足io.Writer接口的约定，所以可以在print中使用它

	fmt.Fprintf(&c, "hello %s", "world")
	fmt.Println(c) //12

	s := ByteString([]byte{'c', 'b', 'd'})
	fmt.Println(s) //{c,b,d,}

}
