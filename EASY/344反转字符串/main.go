package main

func reverseString(s []byte) {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
	println(string(s))
}

func main() {
	s1 := []byte{'h', 's', 't', 'r'}
	reverseString(s1)
}
