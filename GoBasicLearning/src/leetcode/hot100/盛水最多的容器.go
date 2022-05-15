package hot100

/*
 给定一个长度为 n 的整数数组height。有n条垂线，第 i 条线的两个端点是(i, 0)和(i, height[i])
 找出其中的两条线，使得它们与x轴共同构成的容器可以容纳最多的水。
 返回容器可以储存的最大水量。
*/

func maxArea(height []int) int {
	result := 0
	for i, j := 0, len(height)-1; i < j; {
		cur := minInt(height[i], height[j]) * (j - i)
		if cur > result {
			result = cur
		}
		if height[i] < height[j] {
			i++
		} else {
			j--
		}
	}
	return result
}

func minInt(v1, v2 int) int {
	if v1 < v2 {
		return v1
	} else {
		return v2
	}
}
