package main

import "fmt"

func zhuanqian(prices []int) int {
	n := 0
	minNum := prices[0]
	for i := 0; i < len(prices); i++ {
		if minNum > prices[i] {
			minNum = prices[i]
		}
		if prices[i]-minNum > n {
			n = prices[i] - minNum
		}
	}
	fmt.Println(minNum)
	return n
}

func main() {
	fmt.Println(zhuanqian([]int{7, 1, 5, 3, 6, 4}))
}
