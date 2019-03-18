package main

import (
	"fmt"
	"time"
)

func main(){
	p := fmt.Println

	now := time.Now()  //2019-03-18 10:16:42.3485204 +0800 CST m=+0.004988601
	p(now)

	then := time.Date(2009,1,17,20,22,55,651387237,time.FixedZone("CST",8))
	p(then)//2009-01-17 20:22:55.651387237 +0000 UTC

	p(then.Year())
	p(then.Month())
	p(then.Day())
	p(then.Hour())
	p(then.Minute())
	p(then.Second())
	p(then.Nanosecond())
	p(then.Local()) //2009-01-18 04:22:55.651387237 +0800 CST

	p(then.Weekday()) //Saturday
	p(then.Before(now))  //true  时间比较，当前时间之前
	p(then.After(now))  //false
	p(then.Equal(now)) //false

	diff := now.Sub(then)  //89070h3m52.332422863s 输出差值
	p(diff)

	p(diff.Hours())
	p(diff.Minutes())
	p(diff.Seconds())
	p(diff.Nanoseconds())

	p(then.Add(diff))  //2019-03-18 02:28:23.3210838 +0000 CST 相加

	p(then.Add(-diff)) //1998-11-20 14:16:31.232739474 +0000 CST 相减


}
