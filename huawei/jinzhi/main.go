package main

import (
	"fmt"
	"io"
	"math"
)

func main() {
	var result []int
	for {
		var h string
		_, err := fmt.Scanf("%s", &h)
		if err == io.EOF {
			break
		} else {
			bytes := []byte(h)
			r := 0
			for i := 2; i < len(bytes); i++ {
				var a int
				if bytes[i] > 64 {
					a = int(bytes[i] - 55)
				} else {
					a = int(bytes[i] - 48)
				}
				r += a * int(math.Pow(float64(16), float64(len(bytes)-i-1)))
			}
			result = append(result, r)
		}
	}
	for _, r := range result {
		fmt.Println(r)
	}
}
