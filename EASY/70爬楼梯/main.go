package main

import "fmt"

func f(n int) int {
	if n == 1 {
		return 1
	}
	if n == 2 {
		return 2

	}
	s := 0
	a := 1
	b := 2
	for i := 3; i <= n; i++ {
		s = a + b
		a = b
		b = s

	}
	return s
}

func main() {
	fmt.Println(f(5))
}
