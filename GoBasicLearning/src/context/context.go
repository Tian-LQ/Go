package context

// TODO Context
// 根Context：通过context.Background()创建
// 子Context：context.WithCancel(parentContext)创建
// ctx, cancel := context.WithCancel(context.Background())
// 当前Context被取消时，基于他的子context都会被取消
// 接收取消通知<-ctx.Done()
