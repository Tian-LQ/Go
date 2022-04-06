package sync

import (
	"sync/atomic"
)

type Limiter struct {
	// 当前处理请求的上限
	limit int32
	// 处理请求逻辑
	handler func(req interface{}) interface{}

	cnt int32
}

// Reject bool 返回值表示究竟有没有执行
func (l *Limiter) Reject(req interface{}) (interface{}, bool) {
	//cnt := atomic.LoadInt32(&l.cnt)
	//if cnt >= l.limit {
	//	return nil, false
	//}
	// 进来二话不说，直接+1，代表的意思是，我分配一个位置给你
	cnt := atomic.AddInt32(&l.cnt, 1)
	defer atomic.AddInt32(&l.cnt, -1)
	if cnt >= l.limit {
		return nil, false
	}
	res := l.handler(req)
	return res, true
}
