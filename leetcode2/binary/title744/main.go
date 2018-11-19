package main

import (
	"fmt"
)

/*
// 递归方式
func nextGreatestLetter(letters []byte, target byte) byte {
	if letters[0] > target {
		return letters[0]
	} else if letters[len(letters)-1] < target {
		return letters[0]
	}
	result := binarySearch(0, len(letters)-1, letters, target)
	return letters[result]
}

func binarySearch(head int, tail int, letters []byte, target byte) int {
	middle := (head + tail) / 2

	if letters[middle] > target && letters[middle-1] <= target {
		return middle
	} else if letters[middle] > target {
		return binarySearch(head, middle-1, letters, target)
	} else {
		return binarySearch(middle+1, tail, letters, target)
	}

}
*/

func nextGreatestLetter(letters []byte, target byte) byte {

	l := 0
	r := len(letters)
	for l < r {
		m := l + (r-l)/2
		if letters[m] <= target {
			l = m + 1
		} else {
			r = m
		}
	}
	return letters[l%len(letters)]
}

func main() {
	test := []byte{'c', 'f', 'j'}
	var target byte
	target = 'j'
	fmt.Println(string(nextGreatestLetter(test, target)))
}
