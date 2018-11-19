package main

import (
	"fmt"
	"io"
)

func main() {
	for {
		var input string
		_, err := fmt.Scanf("%s", &input)
		if err == io.EOF {
			break
		} else {
			var yde string
			fmt.Scanf("%s", &yde)

			var result float64
			for index, mr := range input {
				if 'A' <= mr && mr <= 'Z' || 'a' <= mr && mr <= 'z' || '0' <= mr && mr <= '9' {
					if yde[index] == '1' {
						result++
					}
				} else if yde[index] == '0' {
					result++
				}
			}
			result = result / float64(len(input)) * 100

			fmt.Print(fmt.Sprintf("%.2f", result))
			fmt.Println("%")

		}
	}

}
