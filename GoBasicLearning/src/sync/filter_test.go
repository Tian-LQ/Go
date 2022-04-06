package sync

import "sync/atomic"

type filter struct {
	// 处理请求逻辑
	handler func(req interface{}) interface{}
	// 0代表接收，1代表拒绝
	reject int32
}

// bool 返回值代表是否接收处理请求
func (f *filter) Handle(req interface{}) (interface{}, bool) {
	if atomic.LoadInt32(&f.reject) > 0 {
		return nil, false
	}
	return f.handler(req), true
}

// 设置服务拒绝请求
func (f *filter) RejectNewRequest() {
	atomic.StoreInt32(&f.reject, 1)
}

// NoBlock
func (f *filter) HandleNoBlock(req interface{}) (interface{}, bool) {
	if f.reject > 0 {
		return nil, false
	}
	return f.handler(req), true
}

// NoBlock
func (f *filter) RejectNewRequestNoBlock() {
	f.reject = 1
}
