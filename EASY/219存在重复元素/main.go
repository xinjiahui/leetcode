package main

/*
给你一个整数数组 nums 和一个整数 k ，判断数组中是否存在两个 不同的索引 i 和 j ，满足 nums[i] == nums[j] 且 abs(i - j) <= k 。
如果存在，返回 true ；否则，返回 false 。
*/
import (
	"fmt"
)

func chongfu(nums []int, k int) bool {
	pos := make(map[int]int)
	for i, num := range nums {
		fmt.Println(num, pos[num], pos)
		if p, ok := pos[num]; ok && i-p <= k {
			return true
		}
		fmt.Println(pos)
		pos[num] = i

	}
	return false
}
func main() {
	fmt.Println(chongfu([]int{1, 2, 1, 1, 2, 3, 6, 7, 3}, 2))
	//fmt.Println(chongfu([]int{1, 2, 3, 4, 5, 6, 1}, 4))
}
