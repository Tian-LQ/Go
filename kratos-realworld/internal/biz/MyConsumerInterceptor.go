package biz

import (
	"github.com/Shopify/sarama"
)

type MyConsumerInterceptor struct {
	str string
}

// OnConsume 该方法在消息返回给 Consumer 程序之前调用。也就是说在开始正式处理消息之前，拦截器会先拦一道，搞一些事情，之后再返回给你。
// OnConsume sarama关于消费者拦截器仅仅实现了OnConsume方法
func (m *MyConsumerInterceptor) OnConsume(message *sarama.ConsumerMessage) {
	if m.str == "" {
		panic("hey, the interceptor has failed [MyConsumerInterceptor]")
	}
	message.Value = []byte(string(message.Value) + m.str)
}
