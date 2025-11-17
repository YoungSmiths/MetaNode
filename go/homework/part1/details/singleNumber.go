package details

/*
*136. 只出现一次的数字
https://leetcode.cn/problems/single-number/
*/
func SingleNumber(nums []int) int {
	var single = 0
	for _, value := range nums {
		single ^= value
	}
	return single
}

func SingleNumber2(nums []int) int {
	var single = 0
	for i := 0; i < len(nums); i++ {
		single ^= nums[i]
	}
	return single
}

func SingleNumber3(nums []int) int {
	var single = 0
	var index = 0
	for {
		if index >= len(nums) {
			break
		}
		single ^= nums[index]
		index++
	}
	return single
}
