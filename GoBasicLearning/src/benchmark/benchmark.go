package benchmark

// TODO Benchmark

// Benchmark用于对代码片段或者第三方库进行性能测试
// 1.方法名以 Benchmark 开头
// 2.参数类型是 *testing.B
// 3.用 b.ResetTimer( ) 和 b.StopTimer() 来隔离与性能测试无关的代码
// 4.性能测试交给 framework 来做，将需要测试的代码放在循环中，循环的次数由 framework 来返回

// go test -bench=.				// 测试所有方法
// go test -bench=functionName	// 测试指定方法
// 通常我们可以在命令的最后加上参数-benchmem进而获取到更多测试信息

// Tips
// B/op：表示的是每次操作分配的内存字节数
// allocs/op：表示每次操作分配内存的次数
