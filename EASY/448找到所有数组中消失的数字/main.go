package main

import (
	"fmt"
	"sort"
	"strconv"
)

func findDisappearedNumbers(nums []int) []int {
	var res []int
	sortInts := make([]int, len(nums))
	copy(sortInts, nums)
	sort.Ints(sortInts)
	fmt.Println(sortInts)
	for i := 1; i <= len(nums); i++ {
		a := sort.SearchInts(sortInts, i)
		if 
		fmt.Println(a)
		//if a == nil {
		res = append(res, a)
		//}
	}
	
	return res

}

func main() {
	var nums []int
	nums = []int{1, 2, 3, 4, 4, 5, 6, 7, 8}
	findDisappearedNumbers(nums)
}
