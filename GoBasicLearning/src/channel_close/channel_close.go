package channel_close

// TODO channel的关闭
// 1.向关闭的channel发送数据，会导致panic
// 2.v,ok := <-channel;其中ok为bool值，true表示正常接受，false表示通道关闭
// 3.所有的channel接收者都会在channel关闭时，立刻从阻塞等待中返回且上述ok值为false
