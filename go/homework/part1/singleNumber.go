package main

import "fmt"

/*
*136. 只出现一次的数字
 */
func singleNumber(nums []int) int {
	var single int = 0
	for _, value := range nums {
		single ^= value
	}
	return single

}

func main() {
	nums := []int{4, 1, 2, 1, 2}
	fmt.Println(singleNumber(nums))
}
