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
	sort.Ints(nums)
	ans := make([][]int, 0)

	// 枚举 a
	for first := 0; first < n; first++ {
		// 需要和上一次枚举的数不相同
		if first > 0 && nums[first] == nums[first-1] {
			continue
		}
		// c 对应的指针初始指向数组的最右端
		third := n - 1
		target := -1 * nums[first]
		// 枚举 b
		for second := first + 1; second < n; second++ {
			// 需要和上一次枚举的数不相同
			if second > first+1 && nums[second] == nums[second-1] {
				continue
			}
			// 需要保证 b 的指针在 c 的指针的左侧
			for second < third && nums[second]+nums[third] > target {
				third--
			}
			// 如果指针重合，随着 b 后续的增加
			// 就不会有满足 a+b+c=0 并且 b<c 的 c 了，可以退出循环
			if second == third {
				break
			}
			if nums[second]+nums[third] == target {
				ans = append(ans, []int{nums[first], nums[second], nums[third]})
			}
		}
	}
	return ans
}

func TestThreeSum(t *testing.T) {
	nums := []int{-4, -1, -1, 0, 1, 2}
	t.Log(threeSum(nums))
}

// 最接近三数之和
