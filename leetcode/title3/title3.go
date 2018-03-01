package title3

func lengthOfLongestSubstring(s string) int {
	tmp := make(map[byte]int)
	start := -1
	maxlen := 0

	for i := 0; i < len(s); i++ {
		last, ok := tmp[s[i]]
		if ok {
			// 如果在 start 之后(不包含start) 已经出现过一次了
			if last > start {
				if maxlen < (i - start - 1) {
					maxlen = i - start - 1
				}
				start = last
			}
		}
		// 记录本次出现的位置
		tmp[s[i]] = i
		// 字符串已经结束并且结尾值与最后一个不重复子字符串不冲突
		if i == len(s)-1 && (!ok || last <= start) {
			if maxlen < (i - start) {
				maxlen = i - start
			}
		}
	}

	return maxlen
}
