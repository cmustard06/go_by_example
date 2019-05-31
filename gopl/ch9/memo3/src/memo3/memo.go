package memo3

import "sync"

/*
这一版Get的实现，调用Get的goroutine会两次获取锁：查找阶段获取一次，
如果查找没有返回任 何内容，那么进入更新阶段会再次获取。在这两次获取锁的中间阶段，其它goroutine可以随意使用 cache。
*/

type result struct {
	value interface{}
	err error
}

type Func func(key string)(interface{}, error)

func New(f Func) *Memo{
	return &Memo{f:f,cache:make(map[string]result)}
}

type Memo struct {
	f Func
	mu sync.Mutex
	cache map[string]result
}

//使用Get线程安全
func (memo *Memo)Get(key string)(value interface{},err error){
	//多个goroutine一起查询cache，发现没有值，然后一 起调用f这个慢不拉叽的函数。
	// 在得到结果后，也都会去更新map。其中一个获得的结果会覆盖掉另 一个的结果。理想情况下是应该避免掉多余的工作的
	//比如相同的URL调用一次f函数就行了，后面直接从函数记忆中读
	memo.mu.Lock()
	r,ok := memo.cache[key]
	memo.mu.Unlock()

	if !ok{
		r.value, r.err = memo.f(key)    //这里是io处理，最应该优化的地方
		memo.mu.Lock()
		memo.cache[key]=r
		memo.mu.Unlock()
	}
	return r.value,r.err
}
/*
go test -run=TestConcurrent -race -v memo3

*/