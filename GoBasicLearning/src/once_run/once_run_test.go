package once_run

import (
	"fmt"
	"sync"
	"testing"
)

type Singleton struct {
}

var singleInstance *Singleton
var once = sync.Once{}

func GetSingletonObj() *Singleton {
	once.Do(func() {
		fmt.Println("Create Obj")
		singleInstance = new(Singleton)
	})
	return singleInstance
}

func TestGetSingletonObj(t *testing.T) {
	var wg sync.WaitGroup
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func() {
			obj := GetSingletonObj()
			//fmt.Printf("the obj's address: %x\n", unsafe.Pointer(obj))
			fmt.Printf("the obj's address: %p\n", obj)
			wg.Done()
		}()
	}
	wg.Wait()
}
