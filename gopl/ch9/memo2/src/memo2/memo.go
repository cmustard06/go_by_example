package memo2

import "sync"

//在memo2中将使用基于监控的同步是cache并发安全
//这需要在Memo加上一个互斥锁，就可以让cache的操作发生在临界区了

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
	memo.mu.Lock()
	defer memo.mu.Unlock()
	r,ok := memo.cache[key]
	if !ok{
		r.value, r.err = memo.f(key)    //这里是io处理，最应该优化的地方，现在由于锁的存在，变成串行运行的了
		memo.cache[key]=r
	}
	return r.value,r.err
}
/*
go test -run=TestConcurrent -race -v memo2
虽然现在是并发安全的了，但是每次对f的调用期间都会持有锁，Get将本来可以并行运行的I/O操作串行化了
本来f是个io操作，访问目标网页，是最应该优化的地方，但是由于锁的存在，这里变成串行的了
因此需要改进
*/