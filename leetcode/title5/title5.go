package title5

func longestPalindrome(s string) string {
	if s == "" || len(s) == 1 {
		return s
	}
	max_start := 0
	max_end := 0

	for i := 0; i < len(s); i++ {
		start := i
		end := i
		// 忽略掉和自己重复的字符
		for end+1 < len(s) && s[end] == s[end+1] {
			end++
			i++
		}
		// 查看两边的字符是否对称相等
		for end < len(s) && start > -1 && s[start] == s[end] {
			start--
			end++
		}
		// 更新最值
		if (end - start - 2) > (max_end - max_start) {
			max_start = start + 1
			max_end = end - 1
		}
	}
	fmt.Println(max_start, max_end)
	return s[max_start : max_end+1]
}
