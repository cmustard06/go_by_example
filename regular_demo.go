package main

import (
	"regexp"
	"fmt"
	"bytes"
)

func main(){
	match,_ := regexp.MatchString("p([a-z]+)ch", "peach")
	fmt.Println(match)

	r,_ := regexp.Compile("p([a-z]+)ch")

	fmt.Println(r.MatchString("peach"))

	fmt.Println(r.FindString("peach punch")) //找到第一个输出

	fmt.Println(r.FindStringIndex("peach punch")) // [0 5]

	fmt.Println(r.FindStringSubmatch("peach punch")) //[peach ea]
	fmt.Println(r.FindStringSubmatchIndex("peach punch")) //[0 5 1 3]

	fmt.Println(r.FindAllString("peach punch pinch",-1))  //-1表示寻找所有

	fmt.Println(r.Match([]byte("peach")))

	//如果想将正则表达式设为常量时
	r = regexp.MustCompile("p([a-z]+)ch")
	fmt.Println(r)

	fmt.Println(r.ReplaceAllString("a peach","<fruit>")) //a <fruit>

	in := []byte("a peach")
	out := r.ReplaceAllFunc(in, bytes.ToUpper)
	fmt.Printf("%s",out)


}
