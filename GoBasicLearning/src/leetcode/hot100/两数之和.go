package hot100

/*
 给定一个整数数组nums和一个整数目标值target，请你在该数组中找出
 和为目标值target的那两个整数，并返回它们的数组下标。
*/

func twoSum(nums []int, target int) []int {
	m := make(map[int]int)
	for index, num := range nums {
		val, ok := m[target-num]
		if ok {
			return []int{val, index}
		} else {
			m[num] = index
		}
	}
	return []int{}
}
