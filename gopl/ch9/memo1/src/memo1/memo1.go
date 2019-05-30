package memo1

/*
函数记忆问题，即缓存函数的结果，达到多次调用但只需要
计算一次的效果
*/

type Func func(key string) (interface{}, error)

type result struct {
	value interface{}
	err   error
}

//Memo缓存了调用函数的结果
type Memo struct {
	f     Func
	cache map[string]result
}

func New(f Func) *Memo {
	return &Memo{f: f, cache: make(map[string]result)}
}

//主意，这是非并发安全的
func (m *Memo) Get(key string) (interface{}, error) {
	r, ok := m.cache[key]
	if !ok {
		r.value, r.err = m.f(key)
		m.cache[key] = r
	}
	return r.value, r.err
}

//func main(){
//	var httpGetBody= HTTPRequest
//	m := New(httpGetBody)
//	//memotest.Sequential(m)
//
//	//下面这个会出错，由于Get函数是非并发安全的,有时候就会出问题，我们需要用test来进行竞态测试
//	Concurrent(m)
//}
