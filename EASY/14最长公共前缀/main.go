package main

import (
	"fmt"
)

func CommonPerfix(strs []string) string {
	a := strs[0]
	min := len(a)
	for i := 0; i < len(strs); i++ {
		s := strs[i]
		if len(s) == 0 {
			return ""
		}
		j := 0
		for ; j < len(a) && j < len(s); j++ {
			if a[j] != s[j] {
				break
			}
		}
		if j < min {
			fmt.Println(min)
			min = j
		}
	}
	return a[0:min]
}

func main() {
	fmt.Println(CommonPerfix([]string{"hasxsc", "htest", "hastest"}))
}
