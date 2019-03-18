package main

import (
	"crypto/sha1"
	"fmt"
	"crypto/md5"
)

func main(){
	s:="sha1 this string"

	h := sha1.New()
	h.Write([]byte(s))
	//bs := h.Sum(nil)
	bs := h.Sum([]byte("123"))
	fmt.Println(s)
	fmt.Printf("%x\n", bs) //cf23df2207d99a74fbe169e3eba035e633b65d94

	// md5
	m := md5.New()
	m.Write([]byte(s))
	bs = m.Sum(nil)
	fmt.Println(bs) //[192 176 107 204 32 135 185 59 59 231 30 28 26 82 76 226]
	b := fmt.Sprintf("%x",bs) //c0b06bcc2087b93b3be71e1c1a524ce2
	fmt.Println(b)

}
