package package_init

// TODO package当中的init方法
// 在main被执行前，所有依赖的package的init方法都会被执行
// 不同包的init函数按照包导入的依赖关系决定执行顺序
// 每个包可以有多个init函数
// 包的每个源文件也可以有多个init函数，这点比较特殊
