package main

import (
	"fmt"
)

func main() {
	c := isValid("[]{}")
	fmt.Println(c)
}
func isValid(s string) bool {
	s3 := []byte{'(', '{', '[', ')', '}', ']'}
	var s2 []byte
	var j int
	j = 0
	if s[0] != s3[0] && s[0] != s3[1] && s[0] != s3[2] {
		return false
	}
	if len(s)%2 != 0 {
		return false
	}
	for i := 0; i < len(s); i++ {
		if s[i] == s3[0] || s[i] == s3[1] || s[i] == s3[2] {
			s2 = append(s2, s[i])
			j++
		} else if s[i] == s3[3] && s2[j-1] == s3[0] && j > 0 {
			j--
			s2 = s2[0:j]
		} else if s[i] == s3[4] && s2[j-1] == s3[1] && j > 0 {
			j--
			s2 = s2[0:j]
		} else if s[i] == s3[5] && s2[j-1] == s3[2] && j > 0 {
			j--
			s2 = s2[0:j]
		} else {
			return false
		}
	}
	if j == 0 {
		return true
	}
	return false
}
