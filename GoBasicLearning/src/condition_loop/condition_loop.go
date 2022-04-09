package condition_loop

// TODO 循环
// Go语言仅仅支持循环关键字for
// for i := 0; i < n; i++

// TODO 条件
// if condition {} else {}
// condition表达式结果必须为布尔值
// 支持变量赋值
// if var declaration; condition {}

// TODO switch条件
// 1.匹配表达式不限制为常量或者整数
// 2.单个case中，可以出现多个结果选项，使用逗号分隔
// 3.在case后无需手动break退出
// 4.可以不设定switch后的匹配表达式，此时整个switch结构与多个if...else..的逻辑作用等同
// 5.switch表达式的结果类型与case表达式的结果类型必须可比较
// 6.如果case表达式的结果值是无类型常量，那么它的类型会被自动地转换为switch表达式的结果类型
// 7.switch语句不允许case表达式中的子表达式结果值存在相等的情况(只针对结果值为常量的情况)
