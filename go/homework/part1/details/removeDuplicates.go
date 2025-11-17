package details

/*
* 26. 删除有序数组中的重复项
https://leetcode.cn/problems/remove-duplicates-from-sorted-array/
*/
func RemoveDuplicates(nums []int) int {
	length := len(nums)
	if length == 0 {
		return 0
	}
	fast := 1
	slow := 1
	for fast < length {
		if nums[fast] != nums[fast-1] {
			nums[slow] = nums[fast]
			slow++
		}
		fast++
	}
	return slow
}
