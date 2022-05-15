package hot100

/*
 给你一个字符串s，找到s中最长的回文子串
*/

func longestPalindrome(s string) string {
	n := len(s)
	if n < 2 {
		return s
	}
	maxLen := 1
	begin := 0
	dp := make([n][n]int, 0)
	for i := 0; i < n; i++ {
		dp[i][i] = true
	}
}
