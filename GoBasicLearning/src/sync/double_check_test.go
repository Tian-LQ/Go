package sync

import (
	"fmt"
	"sync"
	"testing"
)

type SafeMap struct {
	m     map[string]interface{}
	mutex sync.RWMutex
}

// LoadOrStore loaded: 表示返回的是 old_value , 还是 new_value
func (s *SafeMap) LoadOrStore(key string, newValue interface{}) (val interface{}, loaded bool) {
	s.mutex.RLock()
	val, ok := s.m[key]
	s.mutex.RUnlock()
	if ok {
		return val, true
	}
	s.mutex.Lock()
	defer s.mutex.Unlock()
	s.m[key] = newValue
	return newValue, false
}

// LoadOrStore: Double Check(Thread Safety)
func (s *SafeMap) LoadOrStoreDoubleCheck(key string, newValue interface{}) (val interface{}, loaded bool) {
	s.mutex.RLock()
	val, ok := s.m[key]
	s.mutex.RUnlock()
	if ok {
		return val, true
	}
	s.mutex.Lock()
	defer s.mutex.Unlock()
	val, ok = s.m[key]
	// double check
	if ok {
		return val, true
	}
	s.m[key] = newValue
	return newValue, false
}

func TestLoadOrStoreDoubleCheck(t *testing.T) {
	m := &SafeMap{
		m:     map[string]interface{}{},
		mutex: sync.RWMutex{},
	}

	wg := sync.WaitGroup{}
	wg.Add(10)
	for i := 0; i < 10; i++ {
		go func() {
			defer wg.Done()
			value := "Tom"
			newValue, loaded := m.LoadOrStoreDoubleCheck("name", value)
			fmt.Printf("name: %s, loaded: %v\n", newValue.(string), loaded)
		}()
	}
	wg.Wait()
}

type valProvider func() interface{}

type BigObject struct {
	name string
	age  int
}

func GetBigObjectInstance() interface{} {
	return &BigObject{
		name: "tianlq",
		age:  25,
	}
}

// LoadOrStore: Double Check(Thread Safety): Lazy Loading
func (s *SafeMap) LoadOrStoreHeavy(key string, p valProvider) (val interface{}, loaded bool) {
	s.mutex.RLock()
	val, ok := s.m[key]
	s.mutex.RUnlock()
	if ok {
		return val, true
	}
	s.mutex.Lock()
	defer s.mutex.Unlock()
	val, ok = s.m[key]
	// double check
	if ok {
		return val, true
	}
	newValue := p()
	s.m[key] = newValue
	return newValue, false
}

func TestLoadOrStoreHeavy(t *testing.T) {
	m := &SafeMap{
		m:     map[string]interface{}{},
		mutex: sync.RWMutex{},
	}
	wg := sync.WaitGroup{}
	wg.Add(10)
	for i := 0; i < 10; i++ {
		go func() {
			defer wg.Done()
			newValue, loaded := m.LoadOrStoreHeavy("first", GetBigObjectInstance)
			fmt.Printf("first: %v, loaded: %v\n", *newValue.(*BigObject), loaded)
		}()
	}
	wg.Wait()
}
