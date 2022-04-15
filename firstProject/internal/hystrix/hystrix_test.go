package hystrix

import (
	"fmt"
	"github.com/afex/hystrix-go/hystrix"
	"net/http"
	"testing"
	"time"
)

func TestHystrix(t *testing.T) {
	// 配置熔断器相关参数
	hystrix.ConfigureCommand("Mr.T", hystrix.CommandConfig{
		Timeout:                int(3 * time.Second), // 执行 command 的超时时间
		MaxConcurrentRequests:  10,                   // command 的最大并发量
		SleepWindow:            5000,                 // 当熔断器被打开后，SleepWindow 的时间就是控制过多久后去尝试服务是否可用了
		RequestVolumeThreshold: 10,                   // 一个统计窗口 10 秒内请求数量。达到这个请求数量后才去判断是否要开启熔断
		ErrorPercentThreshold:  30,                   // 错误百分比，请求数量大于等于 RequestVolumeThreshold 并且错误率到达这个百分比后就会启动熔断
	})
	_ = hystrix.Do("Mr.T", func() error {
		// talk to other services
		_, err := http.Get("https://www.baidu.com/")
		if err != nil {
			fmt.Printf("get error: %v\n", err)
			return err
		}
		return nil
	}, func(err error) error {
		fmt.Printf("handle error: %v\n", err)
		return nil
	})
}

func TestName(t *testing.T) {
	fmt.Println(time.Now().Unix())
}
