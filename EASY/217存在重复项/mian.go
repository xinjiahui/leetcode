package main

import (
	"fmt"
	"sort"
)

func chongfu(nums []int) bool {
	n := false
	sort.Ints(nums)
	for i := 0; i < len(nums)-1; i++ {
		if nums[i] == nums[i+1] {
			n = true
			break
		}

	}
	return n

}

func main() {
	fmt.Println(chongfu([]int{1, 2, 3, 4, 5, 6, 7}))
	fmt.Println(chongfu([]int{1, 2, 3, 4, 5, 6, 1}))
}
