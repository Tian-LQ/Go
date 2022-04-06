package sync

import "sync"

// 尽量将锁和资源定义在一个结构体里(资源name小写)
type safeResource struct {
	resource interface{}
	lock     sync.Mutex
}

func (s *safeResource) DoSomethingToResource() {
	s.lock.Lock()
	defer s.lock.Unlock()
	// do sth
}

// 拓展阅读：装饰器模式(Decorator)
