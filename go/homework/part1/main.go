package main

import (
	"fmt"

	"github.com/YoungSmiths/go/homework/part1/details"
)

func main() {

	nums := []int{4, 1, 2, 1, 2}
	fmt.Println(details.SingleNumber(nums))
	fmt.Println(details.SingleNumber2(nums))
	fmt.Println(details.SingleNumber3(nums))

	fmt.Println(details.IsPalindrome(121))

	fmt.Println(details.IsValid("[[]{}]"))

	fmt.Println(details.LongestCommonPrefix([]string{"flower", "flow", "flight"}))

	fmt.Println(details.PlusOne([]int{1, 2, 3}))

	fmt.Println(details.RemoveDuplicates([]int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4}))

	fmt.Println(details.Merge([][]int{{1, 3}, {2, 6}, {8, 10}, {15, 18}}))

	fmt.Println(details.TwoSum([]int{21, 7, 2, 15}, 9))
}
