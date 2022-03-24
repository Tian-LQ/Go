package array_slice

// TODO 数组的声明
// var a [3]int					// 声明并初始化为默认零值
// b := [3]int{1,2,3}			// 声明同时初始化
// c := [2][2]int{{1,2},{3,4}}	// 多维数组初始化

// TODO 子切片
// a := [...]int{1,2,3,4,5}
// b := a[start:end]			// b表示的是数组a当中下标[start,end)之间的子切片

// TODO 切片
// var slice []int
// slice := []int{}				// 不推荐这种初始化方法，建议改用make([]int, 0)
// slice := []int{1,2,3,4}
// slice := make([]int, len, cap)
// TODO 切片之间无法判等,切片仅可与nil判等
