package main

import (
	"fmt"
)

func main() {
	var h string
	fmt.Scanf("%s", &h)
	bs := []byte(h)
	m := make(map[byte]bool)
	var result []byte
	for i := len(bs) - 1; i >= 0; i-- {
		b := bs[i]
		if _, ok := m[b]; !ok {
			m[b] = true
			result = append(result, b)
		}
	}

	fmt.Println(string(result))

}
