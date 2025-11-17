package main

import "fmt"

func main() {

	nums := []int{4, 1, 2, 1, 2}
	fmt.Println(singleNumber(nums))
	fmt.Println(singleNumber2(nums))
	fmt.Println(singleNumber3(nums))

	fmt.Println(isPalindrome(121))

	fmt.Println(isValid("[[]{}]"))

	fmt.Println(longestCommonPrefix([]string{"flower", "flow", "flight"}))

	fmt.Println(plusOne([]int{1, 2, 3}))
}
