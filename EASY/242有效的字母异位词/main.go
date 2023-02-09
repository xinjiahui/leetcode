package main

import (
	"fmt"
	"strings"
)

//给定两个字符串 s 和 t ，编写一个函数来判断 t 是否是 s 的字母异位词。

//注意：若 s 和 t 中每个字符出现的次数都相同，则称 s 和 t 互为字母异位词。
func main() {
	s := "aacc"
	t := "ccac"
	s1 := make(map[string]int)
	if len(s) == len(t) {
		for i := 0; i < len(s); i++ {
			if _, ok := s1[string(s[i])]; ok {
				continue
			} else {
				s1[string(s[i])] = strings.Count(s, string(s[i]))
			}
			fmt.Println(s1)
		}
		for j := 0; j < len(t); j++ {
			if _, ok := s1[string(t[j])]; ok && s1[string(t[j])] == strings.Count(t, string(t[j])) {
				fmt.Println(s1[string(t[j])])
				fmt.Println(strings.Count(t, string(t[j])))
				fmt.Println(1)
				continue
			} else {
				//return false
				fmt.Println("false1")
				return
			}
		}

	} else {
		//return false
		fmt.Println("false2")
		return
	}
	fmt.Println("true")
}
