package sync

import (
	"fmt"
	"sync"
	"testing"
)

type Single interface {
	Single()
}

// singleton 结构体私有(防止用户重复创建实例)
type singletonLazyLoading struct {
}

func (s *singletonLazyLoading) Single() {
	fmt.Println("I am single. [lazy loading]")
}

var instanceSingletonLazyLoading *singletonLazyLoading
var instanceOnce sync.Once

// 获取懒汉式single实例
func GetSingletonLazyLoadingInstance() Single {
	instanceOnce.Do(func() {
		fmt.Println("first init single instance. [lazy loading]")
		instanceSingletonLazyLoading = &singletonLazyLoading{}
	})
	return instanceSingletonLazyLoading
}

func TestSingle(t *testing.T) {
	s1 := GetSingletonLazyLoadingInstance()
	s1.Single()
	s2 := GetSingletonLazyLoadingInstance()
	s2.Single()
	if s1 == s2 {
		fmt.Println("s1 == s2")
	}
}
