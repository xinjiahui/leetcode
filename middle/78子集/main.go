package main

import "fmt"

func subsets(nums []int) [][]int {
	var res [][]int
	var tmp []int
	res = append(res, []int{})
	res = append(res, nums)
	for i := 0; i < len(nums); i++ {
		for j := 1; j < len(nums); j++ {
			tmp := []tmp{i, j}
		}

	}
	fmt.Println(res)

	return res

}

func main() {
	nums := []int{1, 2, 3}
	subsets(nums)
}
