package main

import (
	"fmt"
	//	"io"
)

func main() {
	var result []string
	//for {
	for i := 0; i < 2; i++ {
		var h string
		/*
			_, err := fmt.Scanf("%s", &h)
			if err == io.EOF {
				break
			}
		*/
		fmt.Scanf("%s", &h)
		bytes := []byte(h)
		if len(bytes) != 0 {
			i := len(bytes) % 8
			if i != 0 {
				for a := i; a < 8; a++ {
					bytes = append(bytes, '0')
				}
			}
		}

		for i := 0; i < len(bytes)/8; i++ {
			s := string(bytes[8*i : (8 * (i + 1))])
			result = append(result, s)
		}
	}
	for _, r := range result {
		fmt.Println(r)
	}
}
