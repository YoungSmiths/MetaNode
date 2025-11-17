package main

/*
*66. 加一
https://leetcode.cn/problems/plus-one/
*/
func plusOne(digits []int) []int {
	length := len(digits)
	for i := length - 1; i > 0; i -= 1 {
		if digits[i] != 9 {
			digits[i] += 1
			for j := i + 1; j < length; j += 1 {
				digits[j] = 0
			}
			return digits
		}
	}
	ans := make([]int, length+1)
	ans[0] = 1
	return ans
}
