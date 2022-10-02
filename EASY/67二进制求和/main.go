package main

import (
	"fmt"
	"math/big"
)

func sum1(a string, b string) string {
	indexa := len(a) - 1
	indexb := len(b) - 1
	result := ""
	jinwei := 0
	for indexa >= 0 && indexb >= 0 {
		ia := a[indexa] - '0'
		ib := b[indexb] - '0'
		sum := int(ia) + int(ib) + jinwei
		if sum >= 2 {
			jinwei = 1
		} else {
			jinwei = 0
		}
		res := sum%2 + '0'
		result = fmt.Sprintf("%c%s", res, result)
		indexa--
		indexb--

	}
	for indexa >= 0 {
		ia := a[indexa] - '0'
		sum := int(ia) + jinwei
		if sum >= 2 {
			jinwei = 1
		} else {
			jinwei = 0
		}
		res := sum%2 + '0'
		result = fmt.Sprintf("%c%s", res, result)
		indexa--

	}
	for indexb >= 0 {
		ib := b[indexb] - '0'
		sum := int(ib) + jinwei
		if sum >= 2 {
			jinwei = 1
		} else {
			jinwei = 0
		}
		res := sum%2 + '0'
		result = fmt.Sprintf("%c%s", res, result)
		indexb--

	}
	if jinwei == 1 {
		result = fmt.Sprintf("1%s", result)
	}
	return result
}

//方法二 math/big包
func sum2(a string, b string) string {
	ai, _ := new(big.Int).SetString(a, 2)
	bi, _ := new(big.Int).SetString(b, 2)
	ai.Add(ai, bi)
	return ai.Text(2)
}

func main() {
	var a = "101"
	var b = "1010100"
	fmt.Println(sum2(a, b))
	fmt.Println(sum1(a, b))
}
