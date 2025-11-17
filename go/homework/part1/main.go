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

	fmt.Println(removeDuplicates([]int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4}))

	fmt.Println(merge([][]int{{1, 3}, {2, 6}, {8, 10}, {15, 18}}))
}
