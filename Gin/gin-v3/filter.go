package gin_v3

import (
	"fmt"
	"time"
)

// 责任链方式实现 AOP

type Filter func(c *Context)

type FilterBuilder func(next Filter) Filter

// 确保 MetricFilterBuilder 是 FilterBuilder 类型
var _ FilterBuilder = MetricFilterBuilder

func MetricFilterBuilder(next Filter) Filter {
	return func(c *Context) {
		// 执行前的时间
		startTime := time.Now().UnixNano()
		next(c)
		// 执行后的时间
		endTime := time.Now().UnixNano()
		fmt.Printf("run time: %d\n", endTime-startTime)
	}
}
