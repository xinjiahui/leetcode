package main

import (
	"fmt"
	"strings"
)

func change(s string) int {
	iv := strings.Count(s, "IV")
	ix := strings.Count(s, "IX")
	xl := strings.Count(s, "XL")
	xc := strings.Count(s, "XC")
	cd := strings.Count(s, "CD")
	cm := strings.Count(s, "CM")
	m := strings.Count(s, "M") - cm
	d := strings.Count(s, "D") - cd
	c := strings.Count(s, "C") - xc - cd - cm
	l := strings.Count(s, "L") - xl
	x := strings.Count(s, "X") - ix - xc - xl
	v := strings.Count(s, "V") - iv
	i := strings.Count(s, "I") - iv - ix
	fmt.Println(xl, v, m, x)
	return (m*1000 + d*500 + c*100 + l*50 + x*10 + v*5 + i*1 + iv*4 + ix*9 + xl*40 + xc*90 + cd*400 + cm*900)
}

func main() {
	fmt.Println(change("MMMXLV"))
}
