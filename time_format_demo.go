package main

import (
	"fmt"
	"time"
)

func main(){
	p := fmt.Println

	t := time.Now()
	p(t.Format(time.RFC1123)) //Mon, 18 Mar 2019 10:37:52 CST,可以相关的RFC标准格式输出输出

	t1, _ := time.Parse(
		time.RFC3339,
		"2012-11-01T22:08:41+00:00",   //必须与layout格式（RFC3339）一样
	)
	p(t1) //0001-01-01 00:00:00 +0000 UTC

	p(t.Format("3:04PM")) //10:41AM
	p(t.Format("Mon Jan _2 15:04:05 2006")) //Mon Mar 18 10:41:33 2019

	p(t.Format("2006-01-02T15:04:05.999999-07:00")) //Mon Mar 18 10:41:33 2019

	form := "3 04 PM"
	t2,_ := time.Parse(form, "8 41 PM")
	p(t2)

	fmt.Printf("%d-%02d-%02dT%02d:%02d:%02d-00:00\n",
		t.Year(), t.Month(), t.Day(),
		t.Hour(), t.Minute(), t.Second()) //2019-03-18T10:43:55-00:00

	ansic := "Mon Jan _2 15:04:05 2006"
	_, e := time.Parse(ansic, "8:41PM")  //parsing time "8:41PM" as "Mon Jan _2 15:04:05 2006": cannot parse "8:41PM" as "Mon"
	p(e)
}
