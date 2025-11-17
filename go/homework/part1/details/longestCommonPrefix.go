package details

/*
*14. 最长公共前缀
https://leetcode.cn/problems/longest-common-prefix/
*/
func LongestCommonPrefix(strs []string) string {
	length := len(strs)
	if length == 0 {
		return ""
	}
	if length == 1 {
		return strs[0]
	}
	prefix := strs[0]
	count := len(strs)
	for i := 1; i < count; i++ {
		prefix = LongestCommonPrefix1(prefix, strs[i])
	}
	return prefix
}

func LongestCommonPrefix1(prefix string, str string) string {
	prefixLen := len(prefix)
	strLen := len(str)
	// prefixLen 和 strLen 中较小的那个
	minLen := prefixLen
	if strLen < minLen {
		minLen = strLen
	}
	index := 0
	for {
		if index >= minLen {
			break
		}
		if prefix[index] != str[index] {
			break
		} else {
			index++
			continue
		}
	}
	return prefix[:index]
}
