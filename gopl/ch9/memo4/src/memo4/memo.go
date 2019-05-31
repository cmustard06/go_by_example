package memo4

import (
	"sync"
	"fmt"
)

/*
这一版的Memo每一个map元素都是指向一个条目的指针。每一个条目包含对函数f 调用结果的内容缓存。
与之前不同的是这次entry还包含了一个叫ready的channel。
在条目的结果 被设置之后，这个channel就会被关闭，以向其它goroutine广播(§8.9)去读取该条目内的结果是安 全的了。
*/

type result struct {
	value interface{}
	err error
}

//新添加一个类型
type entry struct {
	res result
	ready chan struct{}  //当res准备好的时候就关闭这个channel
}

type Func func(key string)(interface{}, error)

func New(f Func) *Memo{
	return &Memo{f:f,cache:make(map[string]*entry)}
}

type Memo struct {
	f Func
	mu sync.Mutex
	cache map[string]*entry
}

//使用Get线程安全
func (memo *Memo)Get(key string)(value interface{},err error){
	memo.mu.Lock()
	e := memo.cache[key]
	if e==nil{
		//说明这是第一次请求
		//当前这个goroutine复制计算值然后广播给ready
		e = &entry{ready:make(chan struct{})}
		memo.cache[key]=e
		memo.mu.Unlock()   //到这里其他goroutine已经可以行动了
		e.res.value, e.res.err = memo.f(key)
		close(e.ready)
	}else{
		//这是重复的，等着第一次访问的结果
		memo.mu.Unlock()
		fmt.Printf("%v,%v",<-e.ready) //通道关闭返回null
	}
	return e.res.value,e.res.err
}
/*
go test -run=TestConcurrent -race -v memo3

*/