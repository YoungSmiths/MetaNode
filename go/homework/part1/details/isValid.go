package details

/*
*20. 有效的括号
https://leetcode.cn/problems/valid-parentheses/
*/
func IsValid(s string) bool {
	length := len(s)
	if length%2 == 1 {
		return false
	}
	mp := map[byte]byte{
		')': '(',
		']': '[',
		'}': '{',
	}
	//mp := map[byte]byte{}
	//mp := make(map[byte]byte)
	//mp['('] = ')'
	//mp['['] = ']'
	//mp['{'] = '}'
	stack := []byte{}
	for i := 0; i < length; i++ {
		if mp[s[i]] > 0 {
			if len(stack) == 0 || stack[len(stack)-1] != mp[s[i]] {
				return false
			}
			stack = stack[:len(stack)-1]
		} else {
			stack = append(stack, s[i])
		}
	}
	return len(stack) == 0
}
