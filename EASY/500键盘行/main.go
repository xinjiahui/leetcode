package main

import (
	"fmt"
	"strings"
)

func keyboard(words []string) []string {
	s1 := [3]string{"qwertyuiop", "asdfghjkl", "zxcvbnm"}
	var str []string
	for j := 0; j < 3; j++ {
		for _, value := range words {
			count := 0
			value_lower := strings.ToLower(value)

			for i := 0; i < len(value); i++ {
				if strings.Contains(s1[j], string(value_lower[i])) {
					count++
				}
				if count == len(value_lower) {
					str = append(str, value)
				}
			}
		}

	}
	return str

}

func main() {
	fmt.Println(keyboard([]string{"Hello", "Alaska", "Dad", "Peace"}))
}
