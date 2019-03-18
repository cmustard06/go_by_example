package main

import (
	"net/url"
	"fmt"
	"net"
)

func main() {
	//s := "http://www.baidu.com"
	s := "postgres://user:pass@host.com:5432/path?k=v&k=123#f"

	u,err := url.Parse(s)
	if err!=nil{
		panic(err)
	}
	fmt.Println(u.Scheme) //postgres
	fmt.Println(u.User) //user:pass

	fmt.Println(u.User.Username()) //user
	p,_ := u.User.Password()
	fmt.Println(p) //pass

	fmt.Println(u.Host) //host.com:5432

	host,port,_ := net.SplitHostPort(u.Host)
	fmt.Print(host, "\t",port,"\n") //host.com	5432

	fmt.Println(u.Path) ///path
	fmt.Println(u.Fragment) //f

	fmt.Println(u.RawQuery) //k=v
	//解析query
	m,_:= url.ParseQuery(u.RawQuery)
	fmt.Println(m) //map[k:[v 123]]


}
