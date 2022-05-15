package hot100

/*
 给定一个字符串s，请你找出其中不含有重复字符的最长子串的长度。
*/

func lengthOfLongestSubstring(s string) int {
	result := 0
	for i := 0; i < len(s); {
		m := make(map[string]int)
		for j := i; j < len(s); j++ {
			if index, ok := m[string(s[j])]; ok {
				i = index + 1
				break
			} else {
				m[string(s[j])] = j
			}
		}
		if result < len(m) {
			result = len(m)
		}
	}
	return result
}
