package main

import (
	"fmt"
	"strings"
)

/**
给定两个字符串 s 和 t ，判断它们是否是同构的。如果 s 中的字符可以按某种映射关系替换得到 t ，那么这两个字符串是同构的。
每个出现的字符都应当映射到另一个字符，同时不改变字符的顺序。不同字符不能映射到同一个字符上，相同字符只能映射到同一个字上，字符可以映射到自己本身。
**/

func main() {
	s := "badc"
	t := "baba"
	var s1 string
	if len(s) == len(t) {
		m1 := make(map[string]string, 10)
		for i := 0; i < len(s); i++ {
			if value, ok := m1[string(s[i])]; ok {
				if string(t[i]) == value {
					continue
				} else {
					fmt.Println("false")
					return
				}
			} else {
				if strings.Contains(s1, string(t[i])) {
					fmt.Println("false")
					return
				} else {

					m1[string(s[i])] = string(t[i])
					s1 = s1 + string(t[i])
				}
			}
		}

		fmt.Println("true")
		return
	}
	fmt.Println("false")
	return
}
