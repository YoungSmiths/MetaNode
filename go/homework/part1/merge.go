package main

import (
	"sort"
)

/*
*
56. 合并区间
https://leetcode.cn/problems/merge-intervals/description/
*/
func merge(in [][]int) [][]int {
	length := len(in)
	if length == 0 {
		return [][]int{}
	}
	sort.Slice(in, func(i, j int) bool {
		return in[i][0] < in[j][0]
	})
	var merged [][]int
	for i := 0; i < length; i++ {
		L := in[i][0]
		R := in[i][1]
		if len(merged) == 0 || merged[len(merged)-1][1] < L {
			merged = append(merged, []int{L, R})
		} else {
			merged[len(merged)-1][1] = max(merged[len(merged)-1][1], R)
		}
	}
	return merged

}
