package main

import (
	"github.com/Shopify/sarama"
)

type MyProducerInterceptor struct {
	str string
}

// OnSend 该方法会在消息发送之前被调用。如果你想在发送之前对消息“美美容”，这个方法是你唯一的机会。
// OnSend sarama关于生产者拦截器仅仅实现了onSend方法
func (m MyProducerInterceptor) OnSend(message *sarama.ProducerMessage) {
	if m.str == "" {
		panic("hey, the interceptor has failed [MyConsumerInterceptor]")
	}
	v, _ := message.Value.Encode()
	message.Value = sarama.StringEncoder(m.str + string(v))
}
