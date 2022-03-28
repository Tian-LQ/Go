package sync_pool

// TODO sync.Pool对象缓存

// TODO Processor内包含：私有对象(协程安全)和共享池(协程不安全)

// TODO sync.Pool对象的取出
// 1.尝试从私有对象获取
// 2.私有对象不存在，尝试从当前Processor的共享池获取
// 3.如果当前Processor共享池也是空的，那么就尝试去其他Processor的共享池获取
// 4.如果所有子池都是空的，最后就用sync.Pool指定的New函数产生一个新的对象返回

// TODO sync.Pool对象的放回
// 1.如果私有对象不存在则保存为私有对象
// 2.如果私有对象存在，放入当前Processor子池的共享池中

// TODO sync.Pool对象的生命周期
// GC会清除sync.pool缓存的对象
// 对象的缓存有效期为下一次GC之前

// TODO sync.Pool总结
// 适合于通过复用，降低复杂对象的创建和GC代价
// 协程安全，会有锁的开销
// 生命周期受GC影响，不适合于做连接池等，需自己管理生命周期的资源的池化
