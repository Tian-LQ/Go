package sync

import (
	"errors"
	"testing"
)

// TODO 通过约定 来使用无锁解决并发线程安全问题并提高性能
// TODO 无锁(线程不安全)
type Registry struct {
	resources map[string]interface{}
}

func (r *Registry) Register(name string, resource interface{}) {
	r.resources[name] = resource
}

func (r *Registry) Get(name string) (interface{}, error) {
	val, ok := r.resources[name]
	if !ok {
		return nil, errors.New("not found from resources")
	}
	return val, nil
}

func TestRegistry(t *testing.T) {
	registry := &Registry{
		resources: map[string]interface{}{},
	}
	// 要求在应用启动前全部注册好
	registry.Register("a", "a-r")
	registry.Register("b", "b-r")
	registry.Register("c", "c-r")
	registry.Register("d", "d-r")
	registry.Register("e", "e-r")
	runApp()
}

func runApp() {
	// TODO 这里对资源进行并发只读操作
}
