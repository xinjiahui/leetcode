package main

/*
给定一个由 整数 组成的 非空 数组所表示的非负整数，在该数的基础上加一。

最高位数字存放在数组的首位， 数组中每个元素只存储单个数字。

你可以假设除了整数 0 之外，这个整数不会以零开头。

示例1：
输入：digits = [1,2,3]
输出：[1,2,4]
解释：输入数组表示数字 123。
示例2:
输入：digits = [9]
输出：[1,0]
解释：输入数组表示数字 10。
*/
import (
	"fmt"
)

func addOne(digits []int) []int {
	res := make([]int, len(digits))
	copy(res, digits)
	for i := len(digits) - 1; i >= 0; i-- {
		if digits[i] == 9 && i == 0 {
			res[0] = 0
			res = append([]int{1}, res...)
			fmt.Println(i)
			return res
		}
		if res[i]+1 == 10 {
			fmt.Println(res, i)
			res[i] = res[i] + 1 - 10
			fmt.Println(res, i)

		} else {
			res[i] = res[i] + 1
			return res
		}
	}
	return res
}

func main() {
	fmt.Println(addOne([]int{9, 9, 9, 9, 9, 9, 9}))
}
