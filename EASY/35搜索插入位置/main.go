package main

/*
给定一个排序数组和一个目标值，在数组中找到目标值，并返回其索引。如果目标值不存在于数组中，返回它将会被按顺序插入的位置。
*/
import "fmt"

func search(nums []int, target int) int {
	i := 0
	a := 0
	for ; i < len(nums)-1; i++ {
		if nums[i] == target {
			a = i
			return a
		}
		if nums[i] < target && target < nums[i+1] {
			a = i + 1
			return a
		}
	}
	if target <= nums[0] {
		a = 0
		return a
	}
	if target <= nums[len(nums)-1] {
		a = len(nums) - 1
		return a
	}
	if target > nums[len(nums)-1] {
		a = len(nums)
		return a
	}

	return a
}

func main() {
	fmt.Println(search([]int{1, 3, 4, 5}, 7))
}
