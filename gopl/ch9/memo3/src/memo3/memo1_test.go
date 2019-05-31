package memo3

import (
	"memotest"
	"testing"
)

var httpGetBody = memotest.HTTPRequest

func Test(t *testing.T) {
	m := New(httpGetBody)
	memotest.Sequential_ex(t, m)
}


func TestConcurrent(t *testing.T) {
	m := New(httpGetBody)
	memotest.Concurrent_ex(t, m)
}

//使用 go test -v memo1
//检查竞争状态
//go test -run=TestConcurrent -race -v gopl.io/ch9/memo1
