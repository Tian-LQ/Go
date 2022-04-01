package array

import (
	"math"
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
