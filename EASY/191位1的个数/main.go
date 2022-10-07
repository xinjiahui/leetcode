package main

import (
	"fmt"
)

var count = 0

func counts(n uint32) int {
	for n > 0 {
		if (n & 1) == 1 {
			count = count + 1
		}
		n = n >> 1
		fmt.Printf("%b\n", n)

	}
	return count

}
func main() {
	fmt.Println(counts(00000000000000000000000010000000))
}
