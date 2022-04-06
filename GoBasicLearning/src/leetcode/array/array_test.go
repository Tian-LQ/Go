package array

import (
	"math"
	"sort"
	"testing"
)

// 两数之和
func twoSum(nums []int, target int) []int {
	m := make(map[int]int)
	for i, num := range nums {
		if val, ok := m[target-num]; ok {
			return []int{val, i}
		} else {
			m[num] = i
		}
	}
	return []int{}
}

func TestTwoSum(t *testing.T) {
	input := []int{2, 7, 11, 15}
	target := 9
	t.Log(twoSum(input, target))
	t.Log(3 / 2)
}

// 寻找两个正序数组的中位数
func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	len1, len2 := len(nums1), len(nums2)
	var middleLeft, middleRight float64 = 0, 0
	middleLeftIndex, middleRightIndex := (len1+len2-1)/2, (len1+len2)/2
	for i, j, count := 0, 0, 0; count <= middleRightIndex; count++ {
		var val1, val2 int
		if i >= len1 {
			val1 = math.MaxInt
		} else {
			val1 = nums1[i]
		}

		if j >= len2 {
			val2 = math.MaxInt
		} else {
			val2 = nums2[j]
		}

		if val1 < val2 {
			i++
		} else {
			j++
		}

		if count == middleLeftIndex {
			if val1 < val2 {
				middleLeft = float64(val1)
			} else {
				middleLeft = float64(val2)
			}
		}

		if count == middleRightIndex {
			if val1 < val2 {
				middleRight = float64(val1)
			} else {
				middleRight = float64(val2)
			}
		}
	}
	return (middleLeft + middleRight) / 2
}

func TestFindMedianSortedArrays(t *testing.T) {
	nums1 := []int{1, 3}
	nums2 := []int{2, 4}
	t.Log(findMedianSortedArrays(nums1, nums2))
}

// 盛水最多的容器
func Min(x, y int) int {
	if x < y {
		return x
	}
	return y
}

func maxArea(height []int) int {
	result := 0
	i, j := 0, len(height)-1
	for i < j {
		current := Min(height[i], height[j]) * (j - i)
		if current > result {
			result = current
		}
		if height[i] < height[j] {
			i++
		} else {
			j--
		}
	}
	return result
}

func TestMaxArea(t *testing.T) {
	height := []int{1, 8, 6, 2, 5, 4, 8, 3, 7}
	t.Log(maxArea(height))
}

// 三数之和
func threeSum(nums []int) [][]int {
	n := len(nums)
	ans := make([][]int, 0)
	// 特判，如果数组为 nil 或者数组长度小于 3
	if nums == nil || n < 3 {
		return ans
	}
	sort.Ints(nums)
	for first := 0; first < n; first++ {
		// 由于数组已经从小到大排序好，因此若第一个数大于零则表示后面不可能有三数之和为0
		if nums[first] > 0 {
			break
		}
		// 第一个数去重
		if first > 0 && nums[first] == nums[first-1] {
			continue
		}
		// 定义second和third双指针
		second := first + 1
		third := n - 1
		for second < third {
			if nums[first]+nums[second]+nums[third] == 0 {
				ans = append(ans, []int{nums[first], nums[second], nums[third]})
				for second < third && nums[second] == nums[second+1] {
					second++
				}
				for second < third && nums[third] == nums[third-1] {
					third--
				}
				second++
				third--
			} else if nums[first]+nums[second]+nums[third] > 0 {
				third--
			} else {
				second++
			}
		}
	}
	return ans
}

func TestThreeSum(t *testing.T) {
	nums := []int{-4, -2, -1, 0, 3, 2, 5, 6}
	t.Log(threeSum(nums))
}

// 最接近三数之和
func threeSumClosest(nums []int, target int) int {
	n := len(nums)
	if nums == nil || n < 3 {
		return 0
	}
	sort.Ints(nums)

	result := nums[0] + nums[1] + nums[2]

	for first := 0; first < n; first++ {
		second := first + 1
		third := n - 1
		for second < third {
			if math.Abs(float64(nums[first]+nums[second]+nums[third]-target)) < math.Abs(float64(result-target)) {
				result = nums[first] + nums[second] + nums[third]
				for second < third && nums[second] == nums[second+1] {
					second++
				}
				for second < third && nums[third] == nums[third-1] {
					third--
				}
			}
			if nums[first]+nums[second]+nums[third] < target {
				second++
			} else {
				third--
			}
		}
	}
	return result
}

func TestThreeSumClosest(t *testing.T) {
	nums := []int{0, 2, 1, -3}
	target := 1
	t.Log(threeSumClosest(nums, target))
}

// 删除有序数组中的重复项
func removeDuplicates(nums []int) int {
	j := 0
	for i := 0; i < len(nums); i++ {
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}
		nums[j] = nums[i]
		j++
	}
	return j
}

func TestRemoveDuplicates(t *testing.T) {
	nums := []int{1, 1, 2}
	t.Log(removeDuplicates(nums))
	t.Log(nums)
}

// 移除元素
func removeElement(nums []int, val int) int {
	j := 0
	for i := 0; i < len(nums); i++ {
		if nums[i] == val {
			continue
		}
		nums[j] = nums[i]
		j++
	}
	return j
}

func TestRemoveElement(t *testing.T) {
	nums := []int{1, 1, 2}
	t.Log(removeElement(nums, 1))
	t.Log(nums)
}

// 下一个排列
func nextPermutation(nums []int) {
	n := len(nums)
	i := n - 2
	for i >= 0 && nums[i] >= nums[i+1] {
		i--
	}
	if i >= 0 {
		j := n - 1
		for j >= 0 && nums[i] >= nums[j] {
			j--
		}
		nums[i], nums[j] = nums[j], nums[i]
	}
	reverse(nums[i+1:])
}

func reverse(a []int) {
	for i, n := 0, len(a); i < n/2; i++ {
		a[i], a[n-1-i] = a[n-1-i], a[i]
	}
}

func TestNextPermutation(t *testing.T) {
	nums := []int{4, 3, 1, 2}
	nextPermutation(nums)
	t.Log(nums)
}
