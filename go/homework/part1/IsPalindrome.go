package main

import (
	"strconv"
)

/*
*9. 回文数
https://leetcode.cn/problems/palindrome-number/
*/
func isPalindrome(x int) bool {
	str := strconv.Itoa(x)
	length := len(str)
	for i := 0; i <= length/2; i++ {
		if str[i] != str[length-i-1] {
			return false
		}
	}
	return true
}
