package main

import "fmt"

func find(nums []int) int {
	for i := 1; i < len(nums); i++ {
		nums[0] ^= nums[i]
	}
	return nums[0]
}

func yihuo(a int, b int) {
	a = a ^ b
	fmt.Println(a)

}
func main() {
	yihuo(1, 2)
	fmt.Println(find([]int{1, 2, 3, 4, 4, 3, 2, 1, 5}))
}
