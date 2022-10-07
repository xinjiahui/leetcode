package main

import (
	"fmt"
	"strings"
)

func length(s string) int {
	str := strings.Split(s, " ")
	s1 := len(str)
	a := 0
	for i := s1 - 1; i > 0; i-- {
		if len(str[i]) == 0 {
			fmt.Println(i)

		} else {
			a = len(str[i])
			break
		}

	}
	return a
}

func main() {
	fmt.Println(length("a"))
}
