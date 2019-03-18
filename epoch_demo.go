package main

import (
	"time"
	"fmt"
)

func main(){
	now := time.Now()
	secs := now.Unix()
	nanos := now.UnixNano()
	fmt.Println(now, secs, nanos)  //2019-03-18 10:32:30.3300427 +0800 CST m=+0.004984801 1552876350 1552876350330042700

	millis := nanos/1000000 //想要得到毫秒，只能手动进行操作
	fmt.Println(millis)  //1552876390240

	fmt.Println(time.Unix(secs, 0)) //可以将秒换成标准Time类型 2019-03-18 10:35:24 +0800 CST

	fmt.Println(time.Unix(0, nanos)) //2019-03-18 10:35:24.4400324 +0800 CST


}
