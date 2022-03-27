package package_init

import (
	"fmt"
	"testing"
)

// 出现了两个方法签名相同的init方法
func init() {
	fmt.Println("init run!")
}

func init() {
	fmt.Println("init run again!")
}

func TestPackageInit(t *testing.T) {
	t.Log("Hello, World!")
}
